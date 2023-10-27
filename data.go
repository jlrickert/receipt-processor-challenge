package main

import "encoding/json"

// go:embded examples/morning-receipt.json
var rawMorningReceipt string

// go:embed examples/simple-receipt.json
var rawSimpleReceipt string

var simpleScript Reciept

var morningReceipt Reciept

func init() {
	json.Unmarshal([]byte(rawSimpleReceipt), &simpleScript)
	json.Unmarshal([]byte(rawMorningReceipt), &morningReceipt)
}
