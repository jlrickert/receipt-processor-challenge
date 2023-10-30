package main

import "testing"

var (
	example1 = Reciept{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
		Total: "35.35",
	}
	example2 = Reciept{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}
)

func TestGetTotal(t *testing.T) {
	cases := []struct {
		input    string
		expected int64
	}{
		{"56.23", 5623},
		{"10.5", 1050},
		{".2", 20},
		{"0.00", 0},
		{"1.00", 100},
	}
	for _, tc := range cases {
		receipt := Reciept{Total: tc.input}
		actual := receipt.GetTotal()
		if actual != tc.expected {
			t.Errorf(
				"Expected receipt.GetTotal(%s) to equal %d. Got %d",
				tc.input, tc.expected, actual,
			)
		}
	}
}

func TestValidate(t *testing.T) {
	invalidTotalExample := targetReceipt.Copy()
	invalidTotalExample.Total = ""

	invalidRetailerExample := targetReceipt.Copy()
	invalidRetailerExample.Retailer = ""

	cases := []struct {
		input    Reciept
		expected bool
	}{
		{targetReceipt, true},
		{cornerMarketReceipt, true},
		{*invalidTotalExample, false},
		{*invalidRetailerExample, false},
	}
	for _, tc := range cases {
		actual := tc.input.Validate()
		if actual != tc.expected {
			t.Errorf(
				"Expected receipt.Validate() to return %v. Got %v for %v",
				tc.expected, actual, tc.input,
			)
		}
	}
}

func TestGetPoints(t *testing.T) {
	cases := []struct {
		input    Reciept
		expected int64
	}{
		{example1, 28},
		{example2, 109},
	}
	for _, tc := range cases {
		actual := tc.input.GetPoints()
		if actual != tc.expected {
			t.Errorf("Expected receipt.GetTotal() to return %d. Got %d for %v",
				tc.expected, actual, tc.input,
			)
		}
	}
}
