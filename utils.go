package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"regexp"
	"time"
)

var (
	UUID_RE       = regexp.MustCompile(`^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$`)
	RETAIL_RE     = regexp.MustCompile(`^\S+$`)
	PRICE_RE      = regexp.MustCompile(`^\d+.\d{2}$`)
	SHORT_DESC_RE = regexp.MustCompile(`^[\w\s\-]+$`)
)

// pseudoUuid generates a UUID. Not sure how secure this but good enough for
// this.
func pseudoUuid() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func isValidDate(value string) bool {
	_, err := time.Parse(time.DateOnly, value)
	return err == nil
}

func mustParseDate(value string) time.Time {
	d, err := time.Parse(time.DateOnly, value)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return d
}

func isValidTime(value string) bool {
	_, err := time.Parse("15:04", value)
	return err == nil
}

func mustParseTime(value string) time.Time {
	t, err := time.Parse("15:04", value)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	return t
}
