package dao

import (
	"block"
	"util"
)

//添加交易
func AddTransaction(transaction *block.Transaction) {
	sql := "insert into tbl_transaction (address, txType, from_key, to_key, record_addr," +
		" nonce, chain_id, timestamp, payload, user_aec_key, record_id, fromSign, toSign)" +
		" values(?,?,?,?,?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sql, transaction.Address, transaction.TxType, transaction.From, transaction.To, transaction.Value,
		transaction.Nonce, transaction.ChainId, transaction.Timestamp,
		transaction.Payload, transaction.InputData, transaction.RecordId, transaction.FromSign, transaction.ToSign)
	if err != nil {
		panic(err)
	}
}

//修改交易的state值
func UpdateTransactionStateAndChainID(transactions []*block.Transaction, id int64) {
	for _, tx := range transactions {
		tx.ChainId = id
		sql := "update tbl_transaction set state = 1, chain_id = ? where address = ?"
		_, err := db.Exec(sql, tx.ChainId, tx.Address)
		util.CheckErr(err)
	}
}

//GetAllTranactionByChainID gets all the tx in the block which height is ChainID
func GetAllTranactionByChainID(id uint32) []*block.Transaction {
	sql := "select * from tbl_transaction where chain_id = ? and state = 1"
	x := ""
	rows, err := db.Query(sql, id)
	util.CheckErr(err)
	txs := []*block.Transaction{}
	for rows.Next() {
		tx := *new(block.Transaction)
		if err = rows.Scan(&tx.Address, &tx.TxType, &tx.From, &tx.To, &tx.Value, &tx.Nonce,
			&tx.ChainId, &tx.Timestamp, &tx.Payload, &tx.InputData, &tx.RecordId, &tx.FromSign,
			&tx.ToSign, &x); err == nil {
			txs = append(txs, &tx)
		} else {
			panic(err)
		}
	}
	return txs
}
