package controller

import (
	"block"
	"dao"
	"fmt"
	"time"
)

var TXs []*block.Transaction
var CurNUM int

//添加一个交易
func AddTransaction(txType int32, doc string, user string, value string, nonce uint64) {
	tx := block.CreateTransaction(int32(txType), user, doc, value, uint64(nonce))
	TXs = append(TXs, tx)
	CurNUM++
	dao.AddTransaction(tx)
	if CurNUM == 10 {
		fmt.Println(" == 10 new b")
		block := block.CreateBlock(BLOCK.Header.Index+1, BLOCK.Hash, time.Now().Unix(), TXs)
		block.AddTransaction(TXs)
		BLOCK = *block
		CurNUM = 0
		dao.SaveBlock(&BLOCK)
		for i := 0; i < len(TXs); i++ {
			dao.UpdateTransactionStateAndChainID(TXs, BLOCK.Header.Index)
		}
		TXs = nil
	}
}

func GetAllTranactionByChainID(id uint32) []*block.Transaction {
	return dao.GetAllTranactionByChainID(id)
}
