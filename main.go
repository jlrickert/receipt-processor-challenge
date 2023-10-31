package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
	router.HandleFunc("/receipts/process", AddReceiptHandle).Methods(http.MethodPost)
	router.HandleFunc("/receipts/{id}/points", GetReceiptPointsHandle).Methods(http.MethodGet)
}

func init() {
	dataLog.Disable()
	databaseLog.Disable()
	schemaLog.Disable()
}

func WriteJson(w http.ResponseWriter, body []byte) (int, error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	return w.Write(body)
}

// AddReceiptHandle handles a POST to add a receipt
//
//   - Payload: Receipt encoded JSON
//   - Result: JSON containing an id for the receipt
//
// An example:
//
//	```json
//	{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
//	```
func AddReceiptHandle(w http.ResponseWriter, r *http.Request) {
	handleErr := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		return
	}

	reqBodyJson, err := io.ReadAll(r.Body)
	if err != nil {
		handleErr(w, r)
		return
	}

	var receipt Receipt
	err = json.Unmarshal(reqBodyJson, &receipt)
	if err != nil {
		handleErr(w, r)
		return
	}

	if !receipt.Validate() {
		handleErr(w, r)
		return
	}

	var resBody struct {
		Id string `json:"id"`
	}
	resBody.Id = pseudoUuid()
	database.AddReceipt(resBody.Id, receipt)

	reqBodyJson, _ = json.Marshal(resBody)
	WriteJson(w, reqBodyJson)
	return
}

// GetReceiptPointsHandle handles a GET request for for the points of an receipt
//
//   - Path: /receipts/{id}/points
//   - Result: JSON containing an id for the receipt
//
// An example:
//
//	 ```sh
//	 curl -X POST \
//	     -H "Content-Type: application/json" \
//	     --data @examples/morning-receipt.json \
//	     localhost:8080/receipts/process
//	 ```
//
//	```json
//	{ "points": "109" }
//	```
func GetReceiptPointsHandle(w http.ResponseWriter, r *http.Request) {
	handleErr := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		return
	}

	id := mux.Vars(r)["id"]
	receipt := database.GetReceipt(id)

	if receipt == nil {
		handleErr(w, r)
		return
	}

	var response struct {
		Points int64 `json:"points"`
	}
	response.Points = receipt.GetPoints()
	resBodyJson, _ := json.Marshal(response)
	WriteJson(w, resBodyJson)
}

func main() {
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
