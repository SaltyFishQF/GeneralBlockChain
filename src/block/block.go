package block

import (
	"block/pb"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gogo/protobuf/proto"
)

func CreateGenesisBlock(time int64, transactions []*blockpb.Transaction) *blockpb.Block {
	b := blockpb.Block{
		Index:        0,
		PreviousHash: "",
		Timestamp:    0,
		Transaction:  nil,
	}
	serialB, _ := proto.Marshal(&b)
	hashB := sha1.Sum(serialB)
	b.Hash = hex.EncodeToString(hashB[:])
	return &b
}

func CreateBlock(index int64, previousHash string,
	time int64, transactions []*blockpb.Transaction) *blockpb.Block {
	b := blockpb.Block{
		Index:        index,
		PreviousHash: previousHash,
		Timestamp:    time,
		Transaction:  transactions,
	}
	serialB, _ := proto.Marshal(&b)
	hashB := sha1.Sum(serialB)
	b.Hash = hex.EncodeToString(hashB[:])
	return &b
}
