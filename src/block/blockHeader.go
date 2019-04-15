package block

import (
	"bytes"
	"encoding/gob"
	"util"
)

type BlockHeader struct {
	Consensus    *Consensus
	Index        int64
	PreviousHash string
	Digest       []byte
	CloudAddress string
	Timestamp    int64
	MerkleRoot   string
	Version      int32
	Size         int64
}

//HashCode returns the hash of the BlockHeader
func (blockHeader *BlockHeader) HashCode() (string, error) {
	ser, err := blockHeader.Serialize()
	hash := util.ToHash(ser)
	return hash, err
}

//Serialize converts the block header to bytes
func (blockHeader *BlockHeader) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(blockHeader)
	return buf.Bytes(), err
}

//CreateBlockHeader creates a new block header and returns it and its hash
func CreateBlockHeader(index int64, phash string, t int64) (*BlockHeader, string) {
	bh := &BlockHeader{
		Index:        index,
		PreviousHash: phash,
		Timestamp:    t,
		Version:      2,
	}
	hash, err := bh.HashCode()
	util.CheckErr(err)
	return bh, hash
}
