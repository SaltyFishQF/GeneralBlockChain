package block

import (
	"block/pb"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gogo/protobuf/proto"
)

type Block struct {
	index        int64
	hash         string
	previousHash string
	transactions []*Transaction
	timestamp    int64
}

func ToProto(block *Block) proto.Message {
	return &blockpb.Block{
		Index:        block.index,
		PreviousHash: block.previousHash,
		Transaction:  []*blockpb.Transaction{},
		Timestamp:    block.timestamp,
	}
}

func CalHash(block *Block) string {
	serialBlock := ToProto(block)
	byteBlock, _ := proto.Marshal(serialBlock)
	hash := sha1.Sum(byteBlock)
	return hex.EncodeToString(hash[:])
}

func CreateBlock(index int64, preHash string, trans []string, timestamp int64) *Block {
	b := Block{
		index, "", preHash, trans, timestamp,
	}
	b.hash = CalHash(&b)
	return &b
}

func CreateGenesisBlock() *Block {
	b := Block{
		0, "", "", nil, 0,
	}
	b.hash = CalHash(&b)
	return &b
}

func GetHash(b *Block) string {
	return b.hash
}

func GetIndex(b *Block) int64 {
	return b.index
}

func GetTime(b *Block) int64 {
	return b.timestamp
}
