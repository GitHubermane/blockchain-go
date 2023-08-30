package main

import (
	"fmt"

	"gihtub.com/Githubermane/go-blockchain/blockchain"
)

func main() {
	bc := blockchain.CreateBlockchain()

	bc.AddBlock("Second block")
	bc.AddBlock("Third block")
	bc.AddBlock("Fourth block")

	for _, b := range bc.Blocks {
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("Prev Hash: %x\n", b.PrevHash)
		fmt.Printf("Time: %s\n", b.Date)
		fmt.Printf("Data: %s\n\n", b.Data)
	}
}
