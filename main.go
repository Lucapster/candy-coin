package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Index        int
	Hash         []byte
	Timestamp    int64
	PrevHash     []byte
	Transactions []Transaction
	Nonce        int
}

type Transaction struct {
	Sender    string
	Recipient string
	Amount    float64
	DigSig    []byte
}

func (block *Block) MineBlock(diff int) {
	check_proof := false
	for nonce := 1; !check_proof; nonce++ {
		hash := string(GenerateHash(block))
		if hash[:diff] == strings.Repeat("0", 64) {
			fmt.Println("Successfully mined block")
		}
	}
}

func GenerateHash(block *Block) []byte {
	// write each piece of data into the buffer then hash it
	innerBuffer := new(bytes.Buffer)
	binary.Write(innerBuffer, binary.BigEndian, int64(block.Index))
	binary.Write(innerBuffer, binary.BigEndian, block.Timestamp)
	innerBuffer.Write(block.PrevHash)
	innerBuffer.WriteString(fmt.Sprintf("%v", block.Transactions))
	// binary.Write(innerBuffer, binary.BigEndian, int64(block.Nonce))

	hash := sha256.Sum256(innerBuffer.Bytes())
	return hash[:]
}

func (bchain *Blockchain) CreateBlock(prevHash []byte) *Block {
	block := &Block{
		Index:        len(bchain.Blocks),
		Hash:         nil,
		Timestamp:    time.Now().Unix(),
		PrevHash:     prevHash,
		Transactions: nil,
		// Nonce:        nil,
	}
	hash := GenerateHash(block)
	block.Hash = hash
	return block
}

func InitBlockchain() *Blockchain {
	// initialize a blockchain with a genesis
	genesis := &Block{
		Index:     0,
		Timestamp: time.Now().Unix(),
		Hash:      []byte(strings.Repeat("0", 64)),
	}
	return &Blockchain{
		Blocks: []*Block{genesis},
	}
}

// func (bchain *Blockchain) AddBlock(block *Block) {
//     block.PrevHash = bchain.Blocks[bchain_length := len(bchain.Blocks)-1].Hash
//     bchain.Blocks = append(bchain.Blocks, block)
//
// }

func (bchain *Blockchain) AddBlock() {
	// calls CreateBlock and appends returned block to blockchain.Blocks
	index := len(bchain.Blocks)
	prevBlock := bchain.Blocks[index-1]
	new := bchain.CreateBlock(prevBlock.Hash)
	bchain.Blocks = append(bchain.Blocks, new)
}

func main() {
	fmt.Println("Initializing Blockchain")
	bc := InitBlockchain()
	bc.AddBlock()
	fmt.Println("Added Block")
	fmt.Println(bc.Blocks)
	bc.AddBlock()
	fmt.Println("Added Block")
	fmt.Println(bc.Blocks)
	bc.AddBlock()
	fmt.Println("Added Block")
	fmt.Println(bc.Blocks)
}
