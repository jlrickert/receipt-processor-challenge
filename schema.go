package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

var (
	schemaLog = newLogger("schema")
)

// Validate the data in the receipt. If Validate returns false then the data
// from other methods may be undefined.
func (receipt *Receipt) Validate() bool {
	if len(receipt.Retailer) <= 0 {
		return false
	}

	if !isValidDate(receipt.PurchaseDate) {
		return false
	}

	if !isValidTime(receipt.PurchaseTime) {
		return false
	}

	for _, item := range receipt.Items {
		if !item.Validate() {
			return false
		}
	}

	if !PRICE_RE.MatchString(receipt.Total) {
		return false
	}

	return true
}

func (receipt *Receipt) GetTotal() int64 {
	val, _ := strconv.ParseFloat(receipt.Total, 64)
	return int64(val * 100)
}

func (receipt *Receipt) Copy() *Receipt {
	r := Receipt{
		Retailer:     receipt.Retailer,
		PurchaseDate: receipt.PurchaseDate,
		PurchaseTime: receipt.PurchaseTime,
		Items:        []Item{},
		Total:        receipt.Total,
	}
	for _, item := range receipt.Items {
		r.Items = append(r.Items, *item.Copy())
	}
	return &r
}

func (item *Item) Validate() bool {
	if !SHORT_DESC_RE.MatchString(item.ShortDescription) {
		return false
	}

	if !PRICE_RE.MatchString(item.Price) {
		return false
	}
	return true
}

func (item *Item) GetPrice() int64 {
	val, _ := strconv.ParseFloat(item.Price, 64)
	return int64(val * 100)
}

func (item *Item) GetShortDesc() string {
	return strings.TrimSpace(item.ShortDescription)
}

func (item *Item) Copy() *Item {
	return &Item{
		ShortDescription: item.ShortDescription,
		Price:            item.Price,
	}
}

// Calculate points based on various criteria
//
//   - One point for every alphanumeric character in the retailer name.
//   - 50 points if the total is a round dollar amount with no cents.
//   - 25 points if the total is a multiple of `0.25`.
//   - 5 points for every two items on the receipt.
//   - If the trimmed length of the item description is a multiple of 3, multiply
//     the price by `0.2` and round up to the nearest integer. The result is the
//     number of points earned.
//   - 6 points if the day in the purchase date is odd.
//   - 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func (receipt *Receipt) GetPoints() int64 {
	var points int64 = 0

	// One point for every alphanumeric character in the retailer name.
	r := regexp.MustCompile(`[a-zA-Z0-9]`)
	matches := r.FindAllString(receipt.Retailer, -1)
	newPoints := int64(len(matches))
	points += newPoints
	schemaLog.Printf("%d points added. Now %d. value: %v; Reason: One point for ever alphanumeric character in the retailer name",
		newPoints, points, receipt.Retailer,
	)

	// 50 points if the total is a round dollar amount with no cents.
	if receipt.GetTotal()%100 == 0 {
		points += 50
		schemaLog.Printf("%d points added. Now %d. value: %v; Reason: 50 points if the total is a round dollar amount with no cents",
			50, points, receipt.GetTotal(),
		)
	}

	// 25 points if the total is a multiple of `0.25`.
	if receipt.GetTotal()%25 == 0 {
		points += 25
		schemaLog.Printf("%d points added. Now %d. value: %v; Reason: 25 points if the total is a multiple of `0.25`",
			25, points, receipt.GetTotal(),
		)
	}

	// 5 points for every two items on the receipt.
	newPoints = int64(len(receipt.Items) / 2 * 5)
	points += newPoints
	schemaLog.Printf("%d points added. Now %d. value: %d items; Reason: 5 points for every two items on the receipt.",
		newPoints, points, len(receipt.Items),
	)

	// If the trimmed length of the item description is a multiple of 3,
	// multiply the price by `0.2` and round up to the nearest integer. The
	// result is the number of points earned.
	for _, item := range receipt.Items {
		if len(item.GetShortDesc())%3 == 0 {
			price := item.GetPrice() / 5
			if price%100 > 0 {
				price = price - price%100 + 100
			}
			newPoints := price / 100
			points += newPoints
			schemaLog.Printf("%d points added. Now %d. value: %v; length: %d; price: %d, Reason: If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.",
				newPoints, points, item.GetShortDesc(), len(item.GetShortDesc()), item.GetPrice(),
			)
		}
	}

	// 6 points if the day in the purchase date is odd.
	date := mustParseDate(receipt.PurchaseDate)
	if date.Day()%2 == 1 {
		points += 6
		schemaLog.Printf("%d points added. Now %d. value: %v; Reason: 6 points if the day in the purchase date is odd.",
			6, points, receipt.PurchaseDate,
		)
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	timeNow := mustParseTime(receipt.PurchaseTime)
	time2pm := mustParseTime("14:00")
	time4pm := mustParseTime("16:00")
	if time2pm.Before(timeNow) && timeNow.Before(time4pm) {
		points += 10
		schemaLog.Printf("%d points added. Now %d. Value: %v; Reason: 10 points if the time of purchase is after 2:00pm and before 4:00pm.",
			10, points, receipt.PurchaseTime,
		)
	}

	return points
}
