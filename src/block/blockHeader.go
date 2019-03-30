package block

import "block/pb"

func CreateBlockHeader(index int64, phash string, t int64) *blockpb.BlockHeader {
	return &blockpb.BlockHeader{
		Index:        index,
		PreviousHash: phash,
		Timestamp:    t,
	}
}
