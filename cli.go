package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Cli struct {
	bc *BlockChain
}

func (c *Cli) Run() {
	c.validateArgs()

	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addBlock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		c.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			os.Exit(1)
		}
		c.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		c.printChain()
	}
}

func (c *Cli) addBlock(data string) {
	c.bc.AddBlock(data)
}

func (c *Cli) printChain() {
	it := c.bc.Iterator()
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

func (c *Cli) validateArgs() {
	if len(os.Args) < 2 {
		c.printUsage()
		os.Exit(1)
	}
}

func (c *Cli) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  addBlock - Add a block to the top of the chain")
	fmt.Println("  printchain - Print all the blocks of the blockchain")
}
