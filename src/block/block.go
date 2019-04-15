package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"time"
	"util"
)

type Block struct {
	Header      *BlockHeader
	Transaction []*Transaction
	Hash        string
}

//HashCode returns the hash of the block
func (block *Block) HashCode() (string, error) {
	ser, err := block.Serialize()
	hash := util.ToHash(ser)
	return hash, err
}

//Serialize converts block to bytes
func (block *Block) Serialize() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(block)
	return buf.Bytes(), err
}

//CreateGenesisBlock creates the genesis block
func CreateGenesisBlock() *Block {
	bh, hash := CreateBlockHeader(0, "", time.Now().Unix())
	b := Block{
		Header:      bh,
		Transaction: nil,
		Hash:        hash,
	}
	return &b
}

//CreateBlock creates a new block
func CreateBlock(index int64, previousHash string,
	time int64, transactions []*Transaction) *Block {
	bh, hash := CreateBlockHeader(index, previousHash, time)
	b := Block{
		Header:      bh,
		Transaction: transactions,
		Hash:        hash,
	}
	return &b
}

//addTransaction adds transaction to block while a block is created
func (block *Block) AddTransaction(txs []*Transaction) error {
	block.Transaction = txs
	s := *new([]string)
	for i := 0; i < len(txs); i++ {
		s = append(s, txs[i].Id)
	}
	block.Header.MerkleRoot = CalMerkleTreeRoot(s)
	return nil
}

//CalMerkleTreeRoot calculates the hash of transactions
func CalMerkleTreeRoot(hashs []string) string {
	s := *new([]string)
	count := len(hashs)
	var i int
	if count == 1 {
		return hashs[0]
	}
	for ; i < count; i += 2 {
		hash := sha256.Sum256([]byte(hashs[i] + hashs[i+1]))
		s = append(s, hex.EncodeToString(hash[:]))
		if count%2 == 1 && i == count-3 {
			s = append(s, hashs[i+2])
			break
		}
	}
	return CalMerkleTreeRoot(s)
}
