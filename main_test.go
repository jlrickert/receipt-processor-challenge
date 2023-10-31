package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddReceiptHandler(t *testing.T) {
	reqBodyJson, _ := json.Marshal(example1)
	reqBody := bytes.NewReader(reqBodyJson)
	r := httptest.NewRequest(http.MethodPost, "/receipts/process", reqBody)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res := w.Result()
	defer res.Body.Close()

	resBodyJson, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Expected no error. Got %v", err)
	}

	var resData struct {
		Id string `json:"id"`
	}
	err = json.Unmarshal(resBodyJson, &resData)
	if err != nil {
		t.Fatalf("Expected no error. Got %v", err)
	}

	if !UUID_RE.MatchString(resData.Id) {
		t.Errorf("Expected POST \"receipts/process\" to return a valid uuid id. Got %s", resData.Id)
	}
}

func TestGetReceiptPointsHandler(t *testing.T) {
	cases := []struct {
		input    string
		expected int64
	}{
		{
			input:    simpleReceiptKey,
			expected: 31,
		},
		{
			input:    morningReceiptKey,
			expected: 15,
		},
	}

	for _, tc := range cases {
		url := fmt.Sprintf("/receipts/%s/points", tc.input)
		r := httptest.NewRequest(http.MethodGet, url, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, r)

		res := w.Result()
		defer res.Body.Close()

		resBodyJson, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Expected no error. Got %v", err)
		}

		var resData struct {
			Points int64 `json:"points"`
		}
		err = json.Unmarshal(resBodyJson, &resData)
		if err != nil {
			t.Fatalf("Expected no error. Got %v", err)
		}

		actual := resData.Points

		if actual != tc.expected {
			t.Errorf("Expected GET %s to return %d points. Got %d.",
				url, tc.expected, actual,
			)
		}
	}
}
