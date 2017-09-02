package main

import (
	"fmt"
	"johncoin"
)

func main() {
	bc := johncoin.NewBlockchain()

	bc.AddBlock("Send 1 JohnCoin to Bob")
	bc.AddBlock("Send 2 more JohnCoin to Bob")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
