package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func handleConn(conn net.Conn) {

	defer conn.Close()

	io.WriteString(conn, "Enter new Balance:")

	scanner := bufio.NewScanner(conn)

	// take in BPM from stdin and add it to blockchain after conducting necessary validation
	go func() {
		for scanner.Scan() {
			balance, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Printf("%v not a number: %v", scanner.Text(), err)
				continue
			}
			newBlock, err := GenerateBlock(Blockchain[len(Blockchain)-1], balance)
			if err != nil {
				log.Println(err)
				continue
			}
			if CheckBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				ReplaceChain(newBlockchain)
			}

			BcServer <- Blockchain
			io.WriteString(conn, "\nEnter new Balance:")
		}
	}()

	// simulate receiving broadcast
	
	}

}

func StartServer() {
	server, err := net.Listen("tcp", ":"+"8000")

	if err != nil {
		log.Fatal(err)
	}
	log.Println("TCP  Server Started on port 8000")

	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
