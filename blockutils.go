package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"sync"
	"strconv"
	"fmt"
	"strings"
)

var mutex = &sync.Mutex{}
var difficulty = 1

func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.Balance) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func GenerateBlock(oldBlock Block, Balance int) (Block, error) {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Balance = Balance
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	
	return newBlock,nil
}

func CheckBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func ReplaceChain(newBlocks []Block) {
	mutex.Lock()
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
	mutex.Unlock()
}

func GenerateGenesisBlock() Block {
	time := time.Now()
	BcServer = make(chan []Block)
	genesisBlock := Block{}
	genesisBlock = Block{0, time.String(), 0, CalculateHash(genesisBlock), "", difficulty, ""}
	return genesisBlock
}