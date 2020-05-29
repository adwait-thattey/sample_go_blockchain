package main

type Block struct {
	Index     int
	Timestamp string
	Balance       int
	Hash      string
	PrevHash  string
	Difficulty int
	Nonce      string
}

var Blockchain []Block

var BcServer chan []Block