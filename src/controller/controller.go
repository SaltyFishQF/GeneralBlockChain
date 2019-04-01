package controller

import (
	"block"
	"block/pb"
	"dao"
)

var TXs []blockpb.Transaction
var CurNUM int

//计算默克尔树根值
func CalMerkleTreeRoot() {

}

//把交易信息封装入区块
func ToBlock() {
	CalMerkleTreeRoot()
	//b := block.CreateBlock()
}

func AddTransaction(txType int32, doc string, user string, value string, nonce uint64) {
	tx := block.CreateTransaction(int32(txType), user, doc, value, uint64(nonce))
	TXs = append(TXs, *tx)
	CurNUM++
	if CurNUM == 10 {
		ToBlock()
	}
	dao.AddTransaction(tx)
}
