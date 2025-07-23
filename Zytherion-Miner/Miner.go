package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Block struct {
	Index        int               `json:"index"`
	Timestamp    string            `json:"timestamp"`
	Data         map[string]string `json:"data"`
	PreviousHash string            `json:"previous_hash"`
	Hash         string            `json:"hash"`
	Nonce        int               `json:"nonce"`
}

type Blockchain struct {
	Chain      []Block
	Difficulty int
	sync.Mutex
}

var blockchain = &Blockchain{
	Difficulty: 5,
}

const blockFolder = "./blockchain_files"

func calculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + fmt.Sprintf("%v", block.Data) + block.PreviousHash + strconv.Itoa(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func mineBlock(block *Block, difficulty int) {
	targetPrefix := strings.Repeat("0", difficulty)

	for {
		hash := calculateHash(*block)
		if strings.HasPrefix(hash, targetPrefix) {
			block.Hash = hash
			break
		}

		block.Nonce++

		if block.Nonce%100000 == 0 {
			fmt.Printf("\033[33mMining... index: %d, nonce: %d, hash: %s\033[0m\n", block.Index, block.Nonce, hash)
		}
	}
}

func saveBlockAsZythFile(block Block) {
	os.MkdirAll(blockFolder, os.ModePerm)

	raw := fmt.Sprintf("block_%d", block.Index)
	hash := sha256.Sum256([]byte(raw))
	encoded := base64.URLEncoding.EncodeToString(hash[:])
	filename := fmt.Sprintf("%s/%s.zyth", blockFolder, encoded[:20])

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(block)
	if err != nil {
		log.Printf("Failed to encode block %d: %v\n", block.Index, err)
		return
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		log.Printf("Failed to write block file %d: %v\n", block.Index, err)
		return
	}

	fmt.Printf("\033[32mBlock %d saved to %s\033[0m\n", block.Index, filename)
}

func createGenesisBlock() {
	fmt.Println("\033[36mNo existing chain found. Creating genesis block...\033[0m")

	genesis := Block{
		Index:        0,
		Timestamp:    time.Now().Format(time.RFC3339),
		Data:         map[string]string{"message": "Genesis Block"},
		PreviousHash: "0",
		Nonce:        0,
	}
	mineBlock(&genesis, blockchain.Difficulty)
	genesis.Hash = calculateHash(genesis)
	blockchain.Chain = append(blockchain.Chain, genesis)
	saveBlockAsZythFile(genesis)
}

func loadLatestBlock() {
	files, err := ioutil.ReadDir(blockFolder)
	if err != nil || len(files) == 0 {
		createGenesisBlock()
		return
	}

	latest := 0
	var latestBlock Block
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".zyth") {
			data, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", blockFolder, file.Name()))
			if err != nil {
				continue
			}
			var blk Block
			decoder := gob.NewDecoder(bytes.NewReader(data))
			err = decoder.Decode(&blk)
			if err != nil {
				continue
			}
			if blk.Index >= latest {
				latest = blk.Index
				latestBlock = blk
			}
		}
	}
	blockchain.Chain = append(blockchain.Chain, latestBlock)
	fmt.Printf("\033[34mLoaded latest block %d from file.\033[0m\n", latest)
}

func continuousMining() {
	for {
		blockchain.Lock()
		prevBlock := blockchain.Chain[len(blockchain.Chain)-1]
		newBlock := Block{
			Index:        len(blockchain.Chain),
			Timestamp:    time.Now().Format(time.RFC3339),
			Data:         map[string]string{"message": "Mined block at " + time.Now().Format(time.RFC1123), "miner": "Hexaforge"},
			PreviousHash: prevBlock.Hash,
			Nonce:        0,
		}
		blockchain.Unlock()

		fmt.Printf("\033[35mMining block %d...\033[0m\n", newBlock.Index)
		mineBlock(&newBlock, blockchain.Difficulty)
		newBlock.Hash = calculateHash(newBlock)
		fmt.Printf("\033[32mBlock %d mined: %s\033[0m\n", newBlock.Index, newBlock.Hash)

		blockchain.Lock()
		blockchain.Chain = append(blockchain.Chain, newBlock)
		blockchain.Unlock()

		saveBlockAsZythFile(newBlock)
	}
}

func main() {
	loadLatestBlock()
	fmt.Println("\033[36mStarting Continuous Mining Mode (PoW, no delay)...\033[0m")
	continuousMining()
}
