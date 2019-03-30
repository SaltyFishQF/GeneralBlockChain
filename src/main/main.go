package main

import (
	"block"
	"block/pb"
	"fmt"
	"time"
)

func main() {
	blocks := []blockpb.Block{*block.CreateBlock(0, "", 0, nil)}
	cur := blocks[0]

	t := []*blockpb.Transaction{block.CreateTransaction("a", "b", 10)}
	newBlock := block.CreateBlock(cur.Index+1, cur.Hash, time.Now().Unix(), t)
	blocks = append(blocks, *newBlock)
	cur = blocks[cur.Index+1]

	t = []*blockpb.Transaction{block.CreateTransaction("b", "a", 1)}
	newBlock = block.CreateBlock(cur.Index+1, cur.Hash, time.Now().Unix(), t)
	blocks = append(blocks, *newBlock)
	cur = blocks[cur.Index+1]

	t = []*blockpb.Transaction{block.CreateTransaction("b", "a", 4)}
	newBlock = block.CreateBlock(cur.Index+1, cur.Hash, time.Now().Unix(), t)
	blocks = append(blocks, *newBlock)
	cur = blocks[cur.Index+1]

	for _, b := range blocks {
		fmt.Printf("The %d block created at\t", b.Index)
		fmt.Print(time.Unix(b.Timestamp, 0))
		fmt.Println("HASH: ", b.Hash)
	}
}
