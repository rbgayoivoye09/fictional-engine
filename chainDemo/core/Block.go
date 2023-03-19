package core

import (
	myHash "mychain/hashdemo"
	"strconv"
	"time"
)

type Block struct {
	Index        int64  //区块编号
	TimeStamp    int64  //区块时间戳
	PreBlockHash string // 上一个区块的 HASH 值
	Hash         string //当前区块的 HASH 值
	Data         string // 当前区块的数据
}

func calculateBlockHash(block Block) string {
	blockData := strconv.FormatInt(block.Index, 36) + strconv.FormatInt(block.TimeStamp, 36) + block.PreBlockHash + block.Data
	hash := myHash.CalculateHash(blockData)
	return hash
}

func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PreBlockHash = preBlock.Hash
	newBlock.TimeStamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateBlockHash(newBlock)
	return newBlock
}

func GenerateGenesisBlock() Block {
	fakeBlock := Block{}
	fakeBlock.Index = -1
	fakeBlock.Hash = ""
	return GenerateNewBlock(fakeBlock, "")
}
