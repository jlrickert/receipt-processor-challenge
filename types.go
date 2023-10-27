package main

type Reciept struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

func (receipt *Reciept) validate() bool {
	return true
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

func (receipt *Item) validate() bool {
	return true
}

func (receipt *Reciept) GetPoints() int64 {
	return 2
}
