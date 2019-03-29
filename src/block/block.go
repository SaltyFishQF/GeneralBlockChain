package block

import (
	"block/pb"
	"crypto/sha1"
	"encoding/hex"
	"github.com/gogo/protobuf/proto"
)

type Block struct {
	index int64
	previousHash string
	transaction  []string
	timestamp int64
}

func CreateBlock(index int64, preHash string, trans []string, timestamp int64) Block{
	b := Block{
		index,preHash,trans,timestamp,
	}
	return b
}

func ToProto(block *Block) proto.Message {
	return &blockpb.Block{
		PreviousHash: block.previousHash,
		Transaction:  block.transaction,
	}
}

func GetHash(block *Block) string{
	serialBlock := ToProto(block)
	byteBlock, _ := proto.Marshal(serialBlock)
	hash := sha1.Sum(byteBlock)
	return hex.EncodeToString(hash[:])
}