package main

import (
	_ "embed"
	"encoding/json"
)

var (
	//go:embed examples/morning-receipt.json
	rawMorningReceipt string

	//go:embed examples/simple-receipt.json
	rawSimpleReceipt string

	//go:embed examples/target.json
	rawTargetReceipt string

	//go:embed examples/corner-market.json
	rawCornerMarket string

	morningReceipt      Receipt
	simpleReceipt       Receipt
	targetReceipt       Receipt
	cornerMarketReceipt Receipt

	dataLog = newLogger("data")
)

func initReceipt(data string, receipt *Receipt) {
	err := json.Unmarshal([]byte(data), &receipt)
	if err != nil {
		dataLog.Fatalln("Error: ", err, data)
	}
	if !receipt.Validate() {
		dataLog.Fatalln("Invalid receipt: ", receipt)
	}
	dataLog.Printf("Receipt read: %v %v", data, receipt)
}

func init() {
	initReceipt(rawMorningReceipt, &morningReceipt)
	initReceipt(rawSimpleReceipt, &simpleReceipt)
	initReceipt(rawCornerMarket, &cornerMarketReceipt)
	initReceipt(rawTargetReceipt, &targetReceipt)
}
