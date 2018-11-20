package main

import "time"

type Block struct {
	Timestamp int64
	PrevBlockHash []byte
	Hash []byte
	Data []byte
	Nonce int
}


func NewBlock(data string, prevBlockHash []byte) *Block{
	block := &Block{
		Timestamp: time.Now().Unix(),
		PrevBlockHash: prevBlockHash,
		Data: []byte(data),
	}
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}