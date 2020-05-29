package main

import (

	"github.com/davecgh/go-spew/spew"
)


func main() {
	genesisBlock := GenerateGenesisBlock()
	spew.Dump(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)
	StartServer()
}