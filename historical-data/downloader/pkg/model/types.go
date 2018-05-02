package model

const (
	NASDAQ Market = iota
	AMEX
	NYSE
)

type Market uint16

type ListedSymbol struct {
	Symbol       string `json:"Symbol"`
	Name         string `json:"Name"`
	LastScale    string `json:"lastScale"`
	MarketCap    string `json:"MarketCap"`
	AdrTso       string `json:"ADR TSO"`
	IpoYear      string `json:"IPOyear"`
	Sector       string `json:"Sector"`
	Industry     string `json:"Industry"`
	SummaryQuote string `json:"Summary Quote"`
}
