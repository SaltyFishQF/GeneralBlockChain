package block

import "block/pb"

func CreateTransaction(from string, to string, account int32) *blockpb.Transaction {
	return &blockpb.Transaction{
		Doc:   from,
		User:  to,
		Value: account,
	}
}
