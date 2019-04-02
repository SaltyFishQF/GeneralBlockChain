package dao

import (
	"block/pb"
	"fmt"
	"mydb"
)

var db = mydb.DB

func GetLastBlockNum() int64 {
	var index int64
	sql := "select (height) from tbl_block limit 1"
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

func SaveBlock(block2 blockpb.Block) {
	sql := "insert into tbl_block values (?,?,?,?,?,?,?,?,?)"
	_, err := db.Exec(sql, block2.Header.Index, block2.Header.PreviousHash,
		block2.Header.Digest, block2.Header.CloudAddress, block2.Header.Timestamp,
		block2.Header.MerkleRoot, block2.Header.Version, block2.Header.Size, block2.Hash)
	if err != nil {
		panic(err)
	}
}
