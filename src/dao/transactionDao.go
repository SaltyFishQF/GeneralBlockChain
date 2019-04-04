package dao

import (
	"block"
	"block/pb"
	"util"
)

//添加交易
func AddTransaction(transaction *blockpb.Transaction) {
	sql := "insert into tbl_transaction (id, txType, doc, user, value, agentOrganization," +
		" nonce, chain_id, timestamp, payload, inputData, record_id, userSign, docSign)" +
		" values(?,?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sql, transaction.Id, transaction.TxType, transaction.Doc, transaction.User, transaction.Value,
		transaction.AgentOrganization, transaction.Nonce, transaction.ChainId, transaction.Timestamp,
		transaction.Payload, transaction.InputData, transaction.RecordId, transaction.UserSign, transaction.DocSign)
	if err != nil {
		panic(err)
	}
}

//修改交易的state值
func UpdateTransactionStateAndChainID(transactions []*blockpb.Transaction, id uint32) {
	for _, tx := range transactions {
		tx.ChainId = id
		sql := "update tbl_transaction set state = 1, chain_id = ? where id = ?"
		db.Exec(sql, tx.ChainId, block.CalTXHash(tx))
	}
}

func GetAllTranactionByChainID(id uint32) []*blockpb.Transaction {
	sql := "select * from tbl_transaction where chain_id = ?"
	x := ""
	rows, err := db.Query(sql, id)
	util.CheckErr(err)
	txs := []*blockpb.Transaction{}
	for rows.Next() {
		tx := *new(blockpb.Transaction)
		if err = rows.Scan(&tx.Id, &tx.TxType, &tx.Doc, &tx.User, &tx.Value, &tx.AgentOrganization, &tx.Nonce,
			&tx.ChainId, &tx.Timestamp, &tx.Payload, &tx.InputData, &tx.RecordId, &tx.UserSign,
			&tx.DocSign, &x); err == nil {
			txs = append(txs, &tx)
		} else {
			panic(err)
		}
	}
	return txs
}
