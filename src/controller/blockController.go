package controller

import (
	"block"
	"dao"
	"fmt"
)

//检查创世块是否存在
func CheckGenesisBlock() {
	if dao.GetLastBlockNum() == -1 {
		fmt.Println("Generate the Genesis Block")
		b := block.CreateGenesisBlock()
		dao.SaveBlock(*b)
	}
}
