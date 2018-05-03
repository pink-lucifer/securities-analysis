package parser

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"testing"
)

func TestWrapListedSymbol(t *testing.T) {
	mkt := model.NASDAQ
	line := []string{"AAPL", "Apple Inc.", "162.32", "823613790160", "n/a", "1980", "Technology", "Computer Manufacturing", "https://www.nasdaq.com/symbol/aapl"}
	symbol, err := WrapListedSymbol(mkt, line)
	if err != nil {
		t.Fail()
	}
	t.Log(symbol)
}

func TestWrapEodPrice(t *testing.T) {
	sym := "AAPL"
	mkt := model.NASDAQ
	line := []string{"1984-09-07", "0.42216", "0.42728", "0.41704", "0.42216", "23314701"}
	price, err := WrapEodPrice(sym, mkt, line)
	if err != nil {
		t.Fail()
	}

	t.Log(price)
}
