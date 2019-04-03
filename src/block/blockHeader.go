package block

import (
	"block/pb"
	"crypto/sha256"
	"encoding/hex"
	"github.com/golang/protobuf/proto"
)

func CreateBlockHeader(index int64, phash string, t int64) *blockpb.BlockHeader {
	return &blockpb.BlockHeader{
		Index:        index,
		PreviousHash: phash,
		Timestamp:    t,
		Version:      2,
	}
}

func CalHeaderHash(header *blockpb.BlockHeader) string {
	serialBH, err := proto.Marshal(header)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(serialBH)
	return hex.EncodeToString(hash[:])
}
