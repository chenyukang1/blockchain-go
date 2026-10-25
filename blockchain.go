package main

import (
	"log"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type BlockChain struct {
	top []byte
	db  *leveldb.DB
}

const (
	BLOCKS_PATH = "/tmp/blocks"
)

func NewBlockChain() *BlockChain {
	db, err := leveldb.OpenFile(BLOCKS_PATH, &opt.Options{})
	if err != nil {
		log.Panic("open db {} fail", BLOCKS_PATH)
	}

	top, err := db.Get([]byte("l"), &opt.ReadOptions{})
	if err != nil {
		log.Panic(err)
	}

	if top == nil {
		tx, err := db.OpenTransaction()
		genesis := NewGenesisBlock()
		err = tx.Put(genesis.Hash, genesis.Serialize(), &opt.WriteOptions{})
		if err != nil {
			log.Panic(err)
		}

		err = tx.Put([]byte("l"), genesis.Hash, &opt.WriteOptions{})
		if err != nil {
			log.Panic(err)
		}
		tx.Commit()
	}

	return &BlockChain{top, db}
}

func (bc *BlockChain) AddBlock(data string) {
	if bc.db == nil {
		log.Panic("not bind db")
	}

	prevBlockHash, err := bc.db.Get([]byte("l"), &opt.ReadOptions{})
	if err != nil {
		log.Panic(err)
	}

	newBlock := NewBlock(data, prevBlockHash)
	tx, err := bc.db.OpenTransaction()
	err = tx.Put(newBlock.Hash, newBlock.Serialize(), &opt.WriteOptions{})
	if err != nil {
		log.Panic(err)
	}

	err = tx.Put([]byte("l"), newBlock.Hash, &opt.WriteOptions{})
	if err != nil {
		log.Panic(err)
	}
	tx.Commit()

	bc.top = newBlock.Hash
}
