package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]

}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
	fmt.Println("prevBlock", prevBlock)
	fmt.Println("new", new)
	fmt.Println("chain blocks", chain.blocks)

}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{Genesis()}}
}

func main() {

	chain := InitBlockchain()

	chain.AddBlock("First Block after Genesis")
	//chain.AddBlock("Second Block after Genesis")
	//chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash :%x\n", block.PrevHash)
		fmt.Printf("Data in Block :%s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)

	}
}

// & returns the memory address of the variable
// * returns the VALUE of the variable
