package main

import (
	"block"
	"fmt"
	"time"
)

func main() {
	genesisBlock := block.CreateBlock(0, []string{"Hello", "Block", "World"},time.Now().Unix())

	block1 := block.CreateBlock(block.GetHash(&genesisBlock), []string{"first block"})
	block2 := block.CreateBlock(block.GetHash(&block1), []string{"second block"})
	block3 := block.CreateBlock(block.GetHash(&block2), []string{"third block"})

	fmt.Println("genesisBlock:", block.GetHash(&genesisBlock))
	fmt.Println("block1      :", block.GetHash(&block1))
	fmt.Println("block2      :", block.GetHash(&block2))
	fmt.Println("block3      :", block.GetHash(&block3))
}
