package dao

import (
	"block"
	"block/pb"
	"fmt"
	"mydb"
)

var db = mydb.DB

func AddTransaction(transaction *blockpb.Transaction) {
	sql := "insert into tbl_transaction (id, txType, doc, user, value, agentOrganization," +
		" nonce, chain_id, timestamp, payload, inputData, record_id, userSign, docSign)" +
		" values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	res, err := db.Exec(sql, transaction.Id, transaction.TxType, transaction.Doc, transaction.User, transaction.Value,
		transaction.AgentOrganization, transaction.Nonce, transaction.ChainId, transaction.Timestamp,
		transaction.Payload, transaction.InputData, transaction.RecordId, transaction.UserSign, transaction.DocSign)
	fmt.Println(res)
	if err != nil {
		panic(err)
	}
}

//检查创世块是否存在
func CheckGenesisBlock() {
	if GetLastBlockNum() == 0 {
		b := block.CreateGenesisBlock()
		SaveBlock(*b)
	}
}

func GetLastBlockNum() int64 {
	sql := "select 'index' from tbl_block limit 1"
	var index int64
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	if err = rows.Scan(&index); err == nil {

	} else {
		panic(err)
	}
	return index
}

func SaveBlock(block2 blockpb.Block) {
	sql := "insert into tbl_block (index, previousHash, digest, cloudAddress, timestamp," +
		"merkleRoot, version, size) values (?,?,?,?,?,?,?,?)"
	db.Exec(sql, block2.Header.Index, block2.Header.PreviousHash, block2.Header.Digest, block2.Header.CloudAddress,
		block2.Header.Timestamp, block2.Header.Version, block2.Header.Size)
}

func UpdateTransactionState(id []string) {
	//for i := range id{
	//
	//}
}
