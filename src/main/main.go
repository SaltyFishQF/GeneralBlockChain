package main

import (
	"block"
	"fmt"
	"time"
)

func main() {
	blocks := []block.Block{*block.CreateGenesisBlock()}
	cur := blocks[0]
	var t time.Time

	for i := 0; i < 20; i++ {
		newBlock := block.CreateBlock(block.GetIndex(&cur)+1, block.GetHash(&cur), nil, time.Now().Unix())
		blocks = append(blocks, *newBlock)
		cur = blocks[block.GetIndex(newBlock)]
		t = time.Unix(block.GetTime(newBlock), 0)
		fmt.Print(t, ": ")
		fmt.Printf("Block %.2d has been created, \nhash is: %s\n\n",
			block.GetIndex(newBlock), block.GetHash(newBlock))
	}
	//genesisBlock := block.CreateBlock(0, []string{"Hello", "Block", "World"},time.Now().Unix())
	//
	//block1 := block.CreateBlock(block.GetHash(&genesisBlock), []string{"first block"})
	//block2 := block.CreateBlock(block.GetHash(&block1), []string{"second block"})
	//block3 := block.CreateBlock(block.GetHash(&block2), []string{"third block"})
	//
	//fmt.Println("genesisBlock:", block.GetHash(&genesisBlock))
	//fmt.Println("block1      :", block.GetHash(&block1))
	//fmt.Println("block2      :", block.GetHash(&block2))
	//fmt.Println("block3      :", block.GetHash(&block3))
}
