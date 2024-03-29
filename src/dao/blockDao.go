package dao

import (
	"block"
	"fmt"
	"mydb"
)

var db = mydb.DB

func GetLastBlockNum() int64 {
	var index int64
	sql := "select (height) from tbl_block order by height desc limit 1"
	stat, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}
	row := stat.QueryRow()
	err = row.Scan(&index)
	if err != nil {
		return -1
	}
	fmt.Println(index)
	return index
}

func SaveBlock(block2 *block.Block) {
	sql := "insert into tbl_block values (?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sql, block2.Header.Index, block2.Header.PreviousHash,
		block2.Header.Digest, block2.Header.CloudAddress, block2.Header.Timestamp,
		block2.Header.MerkleRoot, block2.Header.Version, block2.Header.Size, block2.Hash)
	if err != nil {
		panic(err)
	}
}

func GetLastBlock() *block.Block {
	block2 := *new(block.Block)
	block2.Header = new(block.BlockHeader)
	sql := "select * from tbl_block order by height desc limit 1"
	row := db.QueryRow(sql)

	row.Scan(&block2.Header.Index, &block2.Header.PreviousHash,
		&block2.Header.Digest, &block2.Header.CloudAddress, &block2.Header.Timestamp,
		&block2.Header.MerkleRoot, &block2.Header.Version, &block2.Header.Size, &block2.Hash)
	return &block2
}

func GetAllBlock() []*block.Block {
	blocks := []*block.Block{}
	sql := "select * from tbl_block"
	rows, err := db.Query(sql)
	for rows.Next() {
		block2 := new(block.Block)
		header := new(block.BlockHeader)
		block2.Header = header
		if err = rows.Scan(&block2.Header.Index, &block2.Header.PreviousHash,
			&block2.Header.Digest, &block2.Header.CloudAddress, &block2.Header.Timestamp,
			&block2.Header.MerkleRoot, &block2.Header.Version, &block2.Header.Size, &block2.Hash); err == nil {
			blocks = append(blocks, block2)
		} else {
			panic(err)
		}
	}
	return blocks
}

func GetBlockByHash(h string) *block.Block {
	sql := "select * from tbl_block where hash = ?"
	rows, err := db.Query(sql, h)
	block2 := new(block.Block)
	header := new(block.BlockHeader)
	for rows.Next() {
		block2.Header = header
		if err = rows.Scan(&block2.Header.Index, &block2.Header.PreviousHash,
			&block2.Header.Digest, &block2.Header.CloudAddress, &block2.Header.Timestamp,
			&block2.Header.MerkleRoot, &block2.Header.Version, &block2.Header.Size, &block2.Hash); err == nil {
		} else {
			panic(err)
		}
	}
	return block2
}
