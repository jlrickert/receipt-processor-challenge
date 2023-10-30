package main

type Database struct {
	receipts map[string]Reciept
}

const (
	simpleReceiptKey  = "FC09442B-C532-E7CF-879A-605E837D3709"
	morningReceiptKey = "493AC2FD-22CE-9280-1853-B5C3480C8E92"
)

var (
	database    Database
	databaseLog = newLogger("database")
)

func init() {
	database = Database{
		receipts: map[string]Reciept{},
	}
	database.AddRecipt(morningReceiptKey, morningReceipt)
	database.AddRecipt(simpleReceiptKey, simpleReceipt)
}

func (db *Database) AddRecipt(id string, receipt Reciept) {
	database.receipts[id] = receipt
	databaseLog.Printf("Adding receipt: %v %v", id, receipt)
}

func (db *Database) GetRecipt(id string) *Reciept {
	receipt, ok := database.receipts[id]
	if !ok {
		databaseLog.Printf("No receipt found for id %v", id)
		return nil
	}

	databaseLog.Printf("Receipt found for id %v. Receipt: %v", id, receipt)
	return &receipt
}
