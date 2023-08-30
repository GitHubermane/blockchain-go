package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Index    uint64
	Hash     []byte
	PrevHash []byte
	Data     []byte
	Date     time.Time
}

type Blockchain struct {
	Blocks []*Block
}

func (b *Block) getHash() {
	data := bytes.Join([][]byte{b.Data, b.PrevHash, []byte(b.Date.String())}, []byte{})
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlock *Block) *Block {
	block := &Block{
		Index:    prevBlock.Index + 1,
		Data:     []byte(data),
		PrevHash: prevBlock.Hash,
		Date:     time.Now(),
	}
	block.getHash()
	return block
}

func CreateGenesisBlock() *Block {
	block := &Block{
		Index:    0,
		Data:     []byte("Genesis Block"),
		PrevHash: []byte(""),
		Date:     time.Now(),
	}
	return block
}

func (c *Blockchain) AddBlock(data string) {
	prevBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := NewBlock(data, prevBlock)
	c.Blocks = append(c.Blocks, newBlock)
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{CreateGenesisBlock()},
	}
}
