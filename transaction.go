package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

type TXOutput struct {
	Value        int    // the coins
	ScriptPubKey string // lock script
}

type TXInput struct {
	Txid      []byte // the id of transaction
	Vout      int    // index of output
	ScriptSig string // data to be used in output's ScriptPubKey
}

const subsidy = 1

func NewCoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{Vin: []TXInput{txin}, Vout: []TXOutput{txout}}
	tx.ID = tx.Hash()

	return &tx
}

func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	return encoded.Bytes()
}

func (tx *Transaction) Hash() []byte {
	copy := *tx
	copy.ID = []byte{}
	hash := sha256.Sum256(copy.Serialize())
	return hash[:]
}
