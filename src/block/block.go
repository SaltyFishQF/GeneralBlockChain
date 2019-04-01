package block

import (
	"block/pb"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gogo/protobuf/proto"
	"time"
)

func CreateGenesisBlock() *blockpb.Block {
	bh := CreateBlockHeader(0, "", time.Now().Unix())
	serialB, _ := proto.Marshal(bh)
	hashB := sha1.Sum(serialB)
	b := blockpb.Block{
		Header:      bh,
		Transaction: nil,
		Hash:        hex.EncodeToString(hashB[:]),
	}
	return &b
}

func CreateBlock(index int64, previousHash string,
	time int64, transactions []*blockpb.Transaction) *blockpb.Block {

	bh := CreateBlockHeader(index, previousHash, time)
	serialB, _ := proto.Marshal(bh)
	hashB := sha1.Sum(serialB)
	b := blockpb.Block{
		Header:      bh,
		Transaction: transactions,
		Hash:        hex.EncodeToString(hashB[:]),
	}
	return &b
}
