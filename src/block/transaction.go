package block

import (
	"block/pb"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func CreateTransaction(txType int32, doc string, user string, value string, nonce uint64) *blockpb.Transaction {
	t := time.Now().Unix()
	hash := sha256.New()
	hash.Write([]byte(doc))
	hash.Write([]byte(user))
	hash.Write([]byte(value))
	hash.Write([]byte(strconv.Itoa(int(txType))))
	hash.Write([]byte(strconv.FormatInt(t, 10)))
	hash.Write([]byte(strconv.FormatInt(int64(nonce), 10)))
	h := hash.Sum(nil)
	return &blockpb.Transaction{
		Id:        hex.EncodeToString(h),
		TxType:    txType,
		Doc:       doc,
		User:      user,
		Value:     value,
		Nonce:     nonce,
		Timestamp: t,
	}
}

func CalTXHash(tx *blockpb.Transaction) string {
	hash := sha256.New()
	hash.Write([]byte(tx.Doc))
	hash.Write([]byte(tx.User))
	hash.Write([]byte(tx.Value))
	hash.Write([]byte(strconv.Itoa(int(tx.TxType))))
	hash.Write([]byte(strconv.FormatInt(tx.Timestamp, 10)))
	hash.Write([]byte(strconv.FormatInt(int64(tx.Nonce), 10)))
	h := hash.Sum(nil)
	return hex.EncodeToString(h[:])
}
