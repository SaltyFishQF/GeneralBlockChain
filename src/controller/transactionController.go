package controller

import (
	"block"
	"common"
	"dao"
	"model"
	"time"
	"util"
)

var TXs []*block.Transaction
var CurNUM int

//添加一个交易
//from and to is the public key of user
func AddTransaction(txType int32, from string, to string, value model.MedicalRecord) {
	nonce, err := dao.GetUserTxNum(to)
	util.CheckErr(err)
	tx := block.CreateTransaction(int32(txType), from, to, value, nonce)
	TXs = append(TXs, tx)
	CurNUM++
	dao.AddTransaction(tx)
	dao.SaveMedicalRecord(value)
	if CurNUM >= common.TransactionNum {
		block := block.CreateBlock(BLOCK.Header.Index+1, BLOCK.Hash, time.Now().Unix(), TXs)
		block.AddTransaction(TXs)
		BLOCK = *block
		CurNUM = 0
		dao.SaveBlock(&BLOCK)
		dao.UpdateTransactionStateAndChainID(TXs, BLOCK.Header.Index)
		TXs = nil
	}
}

func GetAllTranactionByChainID(id uint32) []*block.Transaction {
	return dao.GetAllTranactionByChainID(id)
}
