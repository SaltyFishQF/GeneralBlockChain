package controller

import (
	"block"
	"dao"
	"fmt"
	"time"
)

var BLOCK block.Block

//检查创世块是否存在
func CheckGenesisBlock() {
	if dao.GetLastBlockNum() == -1 {
		fmt.Println("Generate the Genesis Block")
		b := block.CreateGenesisBlock()
		dao.SaveBlock(b)
		BLOCK = *b
	} else {
		BLOCK = *dao.GetLastBlock()
	}
}

func PreCreateBlock(transactions []*block.Transaction) *block.Block {
	block := block.CreateBlock(BLOCK.Header.Index+1, BLOCK.Hash, time.Now().Unix(), transactions)
	fmt.Println("BLOCK index", BLOCK.Header.Index+1)
	fmt.Println("BLOCK index", BLOCK.Header.Index)
	return block
}

func GetAllBlock() []*block.Block {
	return dao.GetAllBlock()
}

func GetBlockByHash(h string) *block.Block {
	return dao.GetBlockByHash(h)
}
