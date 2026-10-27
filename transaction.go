package main

import "fmt"

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
	tx.SetID()

	return &tx
}

func (tx *Transaction) SetID() {
}
