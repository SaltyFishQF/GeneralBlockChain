package dao

import "block/pb"

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
func UpdateTransactionState(id []string) {
	//for i := range id{
	//
	//}
}
