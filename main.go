package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("send 1 BTC to cyk")
	bc.AddBlock("send 2 morce BTC to cyk")

	it := bc.Iterator()
	for {
		block := it.Next()
		if block == nil {
			break
		}
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
