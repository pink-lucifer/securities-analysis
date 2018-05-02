package parser

import (
	"../model"
	"strings"
)

const nyseCsvFile = "../../data/listed-mkt/nyse/companylist-nyse.csv"

func ParseNyseSymbol(ch <-chan []string, eofCh <-chan int, listed chan<- *model.ListedSymbol, done chan<- int) {
	var total, processed int
	for {
		select {
		case line := <-ch:
			if len(line)!=0 {
				listed <- &model.ListedSymbol{
					Symbol:       strings.TrimSpace(line[0]),
					Name:         strings.TrimSpace(line[1]),
					LastScale:    strings.TrimSpace(line[2]),
					MarketCap:    strings.TrimSpace(line[3]),
					AdrTso:       strings.TrimSpace(line[4]),
					IpoYear:      strings.TrimSpace(line[5]),
					Sector:       strings.TrimSpace(line[6]),
					Industry:     strings.TrimSpace(line[7]),
					SummaryQuote: strings.TrimSpace(line[8]),
				}
				processed++
			}
		case eof := <-eofCh:
			total = eof
		default:
			if total != 0 && processed != 0 && total == processed {
				done <- total
				return
			}
		}
	}
}
