package blockchain

import (
	"reflect"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

func CreateGenesisBlock() *Block {
	task := &Task{
		ID:      0,
		Content: []byte("Make Genesis ToDo"),
	}
	block := &Block{
		Index:    0,
		Data:     task,
		PrevHash: []byte(""),
		Date:     time.Now(),
	}
	return block
}

func (c *Blockchain) AddBlock(task *Task) {
	prevBlock := c.Blocks[len(c.Blocks)-1]
	newBlock := NewBlock(task, prevBlock)
	c.Blocks = append(c.Blocks, newBlock)
}

func CreateBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{CreateGenesisBlock()},
	}
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		if !reflect.DeepEqual(currentBlock.Hash, currentBlock.getHash()) {
			return false
		}
		if !reflect.DeepEqual(currentBlock.PrevHash, prevBlock.Hash) {
			return false
		}
	}

	return true
}
