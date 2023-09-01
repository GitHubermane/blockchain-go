package main

import (
	"fmt"

	"gihtub.com/Githubermane/go-blockchain/blockchain"
)

func getTask() func(string) *blockchain.Task {
	counter := 0
	return func(str string) *blockchain.Task {
		counter++
		data := []byte(str)
		return &blockchain.Task{
			ID:      uint64(counter),
			Content: data,
		}
	}
}
func main() {
	bc := blockchain.CreateBlockchain()
	task := getTask()
	bc.AddBlock(task("Second block"))
	bc.AddBlock(task("Third block"))
	bc.AddBlock(task("Fourth block"))

	for _, b := range bc.Blocks {
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("Prev Hash: %x\n", b.PrevHash)
		fmt.Printf("Time: %s\n", b.Date)
		fmt.Printf("Data (Task): \n")
		fmt.Printf("\tID: %x\n", b.Data.ID)
		fmt.Printf("\tContent: %s\n\n", b.Data.Content)

	}
}
