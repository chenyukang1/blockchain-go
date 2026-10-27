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

func NewBlockChain(address string) *BlockChain {
	db, err := leveldb.OpenFile(BLOCKS_PATH, &opt.Options{})
	if err != nil {
		log.Panic("open db {} fail", BLOCKS_PATH)
	}

	top, err := db.Get([]byte("l"), &opt.ReadOptions{})
	if err != nil {
		if err == leveldb.ErrNotFound {
			tx, err := db.OpenTransaction()
			cbtx := NewCoinbaseTx(address, "")
			genesis := NewGenesisBlock(cbtx)
			err = tx.Put(genesis.Hash, genesis.Serialize(), &opt.WriteOptions{})
			if err != nil {
				log.Panic(err)
			}

			err = tx.Put([]byte("l"), genesis.Hash, &opt.WriteOptions{})
			if err != nil {
				log.Panic(err)
			}
			tx.Commit()
		} else {
			log.Panic(err)
		}
	}

	return &BlockChain{top, db}
}

func (bc *BlockChain) AddBlock(block *Block) {
	if bc.db == nil {
		log.Panic("not bind db")
	}

	prevBlockHash, err := bc.db.Get([]byte("l"), &opt.ReadOptions{})
	if err != nil {
		log.Panic(err)
	}
	block.PrevBlockHash = prevBlockHash

	tx, err := bc.db.OpenTransaction()
	err = tx.Put(block.Hash, block.Serialize(), &opt.WriteOptions{})
	if err != nil {
		log.Panic(err)
	}

	err = tx.Put([]byte("l"), block.Hash, &opt.WriteOptions{})
	if err != nil {
		log.Panic(err)
	}
	tx.Commit()

	bc.top = block.Hash
}

type BlockChainIterator struct {
	currentHash []byte
	db          *leveldb.DB
}

func (bc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{bc.top, bc.db}
}

func (bci *BlockChainIterator) Next() *Block {
	blockData, err := bci.db.Get([]byte(bci.currentHash), &opt.ReadOptions{})
	if err != nil {
		if err == leveldb.ErrNotFound {
			return nil
		}
		log.Panic(err)
	}

	block := DeserializeBlock(blockData)
	bci.currentHash = block.PrevBlockHash
	return block
}
