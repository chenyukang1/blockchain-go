package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Cli struct {
	bc *BlockChain
}

func (c *Cli) Run() {
	c.validateArgs()

	createBlockchainCmd := flag.NewFlagSet("createBlockchain", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addBlock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	blockchainAddress := createBlockchainCmd.String("address", "", "Blockchain address")
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "createBlockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
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

	if createBlockchainCmd.Parsed() {
		if *blockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		c.createBlockchain(*blockchainAddress)
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

func (c *Cli) createBlockchain(address string) {
	NewBlockChain(address)
}

func (c *Cli) addBlock(data string) {
	// c.bc.AddBlock(data)
}

func (c *Cli) printChain() {
	it := c.bc.Iterator()
	for {
		block := it.Next()
		if block == nil {
			break
		}
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := NewProofOfWork(block)
		fmt.Printf("Proof of Work: %s\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
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
	fmt.Println("  printChain - Print all the blocks of the blockchain")
}
