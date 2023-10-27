package main

import (
	"crypto/rand"
	"fmt"
)

type Database struct {
	receipts map[string]Reciept
}

var database Database

func init() {
	database = Database{
		receipts: map[string]Reciept{
			pseudoUuid(): simpleScript,
			pseudoUuid(): morningReceipt,
		},
	}
}

// Ripped https://stackoverflow.com/questions/15130321/is-there-a-method-to-generate-a-uuid-with-go-language
// TODO: use a real library
func pseudoUuid() (uuid string) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return
}

func (db *Database) AddRecipt(receipt Reciept) string {
	id := pseudoUuid()
	database.receipts[id] = receipt
	return id
}

func (db *Database) GetRecipt(id string) *Reciept {
	receipt, ok := database.receipts[id]
	if !ok {
		return nil
	}
	return &receipt
}
