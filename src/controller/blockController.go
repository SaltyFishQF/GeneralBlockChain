package controller

import (
	"block"
	"block/pb"
	"dao"
	"fmt"
	"time"
)

var BLOCK blockpb.Block

//检查创世块是否存在
func CheckGenesisBlock() {
	if dao.GetLastBlockNum() == -1 {
		fmt.Println("Generate the Genesis Block")
		b := block.CreateGenesisBlock()
		dao.SaveBlock(b)
	} else {
		BLOCK = *dao.GetLastBlock()
	}
}

func PreCreateBlock(transactions []*blockpb.Transaction) *blockpb.Block {
	block := block.CreateBlock(BLOCK.Header.Index+1, BLOCK.Hash, time.Now().Unix(), transactions)
	return block
}

func GetAllBlock() []*blockpb.Block {
	return dao.GetAllBlock()
}

func GetBlockByHash(h string) *blockpb.Block {
	return dao.GetBlockByHash(h)
}
