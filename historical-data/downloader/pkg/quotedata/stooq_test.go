package quotedata

import (
	"../model"
	"testing"
	"time"
)

func TestHistoricalQuoteInitialDownload(t *testing.T) {
	symbol := "APPL"
	err := HistoricalQuoteInitialDownload(model.NASDAQ, symbol, time.Now())
	if err != nil {
		t.Fatal(err)
	}
}
