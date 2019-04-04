package controller

import (
	"block"
	"block/pb"
	"common"
	"crypto/sha256"
	"dao"
	"encoding/hex"
	"fmt"
	"github.com/gogo/protobuf/proto"
)

var TXs []*blockpb.Transaction
var CurNUM int

//计算默克尔树根值
func CalMerkleTreeRoot(hashs []string, count uint) string {
	fmt.Println("===count:", count, "====")
	s := *new([]string)
	var i uint
	if count == 1 {
		fmt.Println(hashs[0])
		return hashs[0]
	}
	for ; i < count; i += 2 {
		fmt.Println("***i:", i, "***")
		hash := sha256.Sum256([]byte(hashs[i] + hashs[i+1]))
		s = append(s, hex.EncodeToString(hash[:]))
		if count%2 == 1 && i == count-3 {
			s = append(s, hashs[i+2])
			fmt.Println("break")
			break
		}
	}
	return CalMerkleTreeRoot(s, (count+1)>>1)
}

//计算默克尔树，预处理
func PreCalMerkleTreeRoot(transactions []*blockpb.Transaction) string {
	s := *new([]string)
	for i := 0; i < common.TransactionNum; i++ {
		serialB, _ := proto.Marshal(transactions[i])
		hash := sha256.Sum256(serialB)
		s = append(s, hex.EncodeToString(hash[:]))
	}
	return CalMerkleTreeRoot(s, common.TransactionNum)
}

//把交易信息封装入区块
func AddTransactionToBlock(transactions []*blockpb.Transaction) {
	block2 := PreCreateBlock(transactions)
	fmt.Println("---index---   =", block2.Header.Index)
	//处理Block Header
	merkleRoot := PreCalMerkleTreeRoot(transactions)
	block2.Header.MerkleRoot = merkleRoot
	block2.Header.Size = common.TransactionNum
	//处理Block
	block2.Hash = block.CalHeaderHash(block2.Header)

	BLOCK = *block2
	dao.SaveBlock(block2)
	dao.UpdateTransactionStateAndChainID(transactions, uint32(BLOCK.Header.Index))
}

//添加一个交易
func AddTransaction(txType int32, doc string, user string, value string, nonce uint64) {
	tx := block.CreateTransaction(int32(txType), user, doc, value, uint64(nonce))
	TXs = append(TXs, tx)
	CurNUM++
	if CurNUM == 10 {
		AddTransactionToBlock(TXs)
	}
	dao.AddTransaction(tx)
}

func GetAllTranactionByChainID(id uint32) []*blockpb.Transaction {
	return dao.GetAllTranactionByChainID(id)
}
