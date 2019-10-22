package main

import(
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


type Mahastudent struct{
	Nama		string
	Jurusan		string
	TahunMasuk	int
}

type Block struct{
	Index		int
	Timestamp	string
	Mahastudent
	Hash		string
	PrevHash	string
}

var BlockChain []Block

func createHash(block Block) string{
	record := string(block.Index) + block.Timestamp + string(block.Mahastudent.TahunMasuk) + block.Mahastudent.Jurusan + block.Mahastudent.Nama + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Mahasiswa Mahastudent)(Block, error){
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Mahastudent = Mahasiswa
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = createHash(newBlock)

	return newBlock, nil
}

func isValidBlock(newBlock, oldBlock Block) bool{
	if oldBlock.Index + 1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if createHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}