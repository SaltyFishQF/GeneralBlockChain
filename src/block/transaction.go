package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"strconv"
	"time"
	"util"
)

type Transaction struct {
	Id                string
	TxType            int32
	From              string
	To                string
	Value             string
	AgentOrganization string
	Nonce             uint64
	ChainId           int64
	Timestamp         int64
	Payload           []byte
	InputData         string
	RecordId          string
	UserSign          []byte
	DocSign           []byte
}

//HashCode returns the hash of transaction
func (tx *Transaction) HashCode() (string, error) {
	ser, err := tx.Serialize()
	hash := util.ToHash(ser)
	return hash, err
}

//Serialize converts Transaction struct to []byte
func (tx *Transaction) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(tx)
	return buf.Bytes(), err
}

//CalHash returns the hash of the information in transaction
func (tx *Transaction) CalHash() string {
	hash := sha256.New()
	hash.Write([]byte(tx.From))
	hash.Write([]byte(tx.To))
	hash.Write([]byte(tx.Value))
	hash.Write([]byte(strconv.Itoa(int(tx.TxType))))
	hash.Write([]byte(strconv.FormatInt(tx.Timestamp, 10)))
	hash.Write([]byte(strconv.FormatInt(int64(tx.Nonce), 10)))
	h := hash.Sum(nil)
	return hex.EncodeToString(h[:])
}

//CreateTransaction creates a new transaction
func CreateTransaction(txType int32, from string, to string, value string, nonce uint64) *Transaction {
	t := time.Now().Unix()
	tx := Transaction{
		TxType:    txType,
		From:      from,
		To:        to,
		Value:     value,
		Nonce:     nonce,
		Timestamp: t,
	}
	tx.Id = tx.CalHash()
	return &tx
}

func UpLoadTransaction() {

}
