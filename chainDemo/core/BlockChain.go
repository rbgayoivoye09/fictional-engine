package core

import (
	"fmt"
	"log"
)

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	GenesisBlock := GenerateGenesisBlock()
	bc := BlockChain{}
	bc.AppendBlock(&GenesisBlock)
	return &bc
}

func (bc *BlockChain) SendData(blockData string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, blockData)
	bc.AppendBlock(&newBlock)
}

func (bc *BlockChain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isVaild(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invaild block")
	}
}
func (bc *BlockChain) Print() {
	bc.prinfBlockChain()
}

func (bc *BlockChain) prinfBlockChain() {
	for _, b := range bc.Blocks {
		fmt.Printf("Index: %d\n", b.Index)
		fmt.Printf("Prev.Hash: %s\n", b.PreBlockHash)
		fmt.Printf("Curr.Hash: %s\n", b.Hash)
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("TimeStamp: %d\n\n", b.TimeStamp)

	}
}

func isVaild(newBlcok, oldBlock Block) bool {
	if newBlcok.Index-1 != oldBlock.Index {
		log.Printf("newBlcok.Index: %v\n", newBlcok.Index)
		log.Printf("oldBlock.Index: %v\n", oldBlock.Index)
		log.Println("newBlcok.Index-1 != oldBlock.Index")
		return false
	}
	if newBlcok.PreBlockHash != oldBlock.Hash {
		log.Println("newBlcok.PreBlockHash != oldBlock.Hash")
		return false
	}
	if calculateBlockHash(newBlcok) != newBlcok.Hash {
		log.Printf("calculateBlockHash(newBlcok) != newBlcok.Hash")
		return false
	}

	return true
}
