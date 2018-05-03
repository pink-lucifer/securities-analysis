package quotedata

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"testing"
	"time"
)

func TestHistoricalQuoteInitialDownload(t *testing.T) {
	symbol := "AAPL"
	err := HistoricalQuoteInitialDownload(model.NASDAQ, symbol, time.Now())
	if err != nil {
		t.Fatal(err)
	}
}
