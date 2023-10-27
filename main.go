package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, body []byte) (int, error) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	return w.Write(body)
}

func AddReceiptHandle(w http.ResponseWriter, r *http.Request) {
	handleErr := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		return
	}

	if r.Method != "POST" {
		w.Header().Add("Allow", "POST")
		handleErr(w, r)
		return
	}

	reqBodyJson, err := io.ReadAll(r.Body)
	if err != nil {
		handleErr(w, r)
		return
	}

	var receipt Reciept
	err = json.Unmarshal(reqBodyJson, &receipt)
	if err != nil {
		handleErr(w, r)
		return
	}

	if !receipt.validate() {
		handleErr(w, r)
		return
	}

	var resBody struct {
		Id string `json:"id"`
	}
	resBody.Id = database.AddRecipt(receipt)

	reqBodyJson, _ = json.Marshal(resBody)
	WriteJson(w, reqBodyJson)
	return
}

func GetReceiptHandle(w http.ResponseWriter, r *http.Request) {
	handleErr := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		return
	}

	if r.Method != "GET" {
		w.Header().Add("Allow", "GET")
		handleErr(w, r)
		return
	}

	id := mux.Vars(r)["id"]
	receipt := database.GetRecipt(id)

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

func handleRequests() {
	router := mux.NewRouter()

	router.HandleFunc("/reciepts/process", AddReceiptHandle)
	router.HandleFunc("/reciepts/{id}/points", GetReceiptHandle)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequests()
}
