package block

import (
	"algorithm"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"model"
	"strconv"
	"time"
	"util"
)

type Transaction struct {
	Address string
	TxType  int32
	From    string
	To      string
	Value   string
	//AgentOrganization string
	Nonce     uint64
	ChainId   int64
	Timestamp int64
	Payload   []byte
	InputData string
	RecordId  string
	FromSign  []byte
	ToSign    []byte
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
func CreateTransaction(txType int32, from string, to string, value model.MedicalRecord, nonce uint64) *Transaction {
	t := time.Now().Unix()

	pt, _ := hex.DecodeString(value.Addr)
	pk, _ := hex.DecodeString(to)
	bvalue, err := algorithm.ECCEncrypt(pt, algorithm.ToECDSAPub(pk))
	util.CheckErr(err)
	svalue := hex.EncodeToString(bvalue)

	tx := Transaction{
		TxType:    txType,
		From:      from,
		To:        to,
		Value:     svalue,
		Nonce:     nonce,
		Timestamp: t,
	}

	tx.Address = tx.CalHash()
	return &tx
}

//todo: To公钥加密Tx的Hash作为Tx的地址， 存入数据库
