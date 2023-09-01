package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"time"
)

type Task struct {
	ID      uint64
	Content []byte
}

type Block struct {
	Index    uint64
	Hash     []byte
	PrevHash []byte
	Data     *Task
	Date     time.Time
}

func (b *Block) getHash() []byte {
	//Converting uint64 to []byte
	taskId := make([]byte, 8)
	blockId := make([]byte, 8)
	binary.LittleEndian.AppendUint64(taskId, b.Data.ID)
	binary.LittleEndian.AppendUint64(blockId, b.Index)

	data := bytes.Join(
		[][]byte{
			taskId,
			blockId,
			b.Data.Content,
			b.PrevHash,
			[]byte(b.Date.String())},
		[]byte{})

	hash := sha256.Sum256(data)
	return hash[:]
}

func NewBlock(task *Task, prevBlock *Block) *Block {
	block := &Block{
		Index:    prevBlock.Index + 1,
		Data:     task,
		PrevHash: prevBlock.Hash,
		Date:     time.Now(),
	}
	block.Hash = block.getHash()
	return block
}
