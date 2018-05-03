package parser

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
)

func ParseAndWrapEodPrice(sym string, mkt model.Market, ch <-chan []string, eofCh <-chan int, listed chan<- *persist.TblEodPrice, done chan<- int) {
	var total, processed int
	for {
		select {
		case line := <-ch:
			if len(line) != 0 {
				price, err := WrapEodPrice(sym, mkt, line)
				if err != nil {
					break
				}
				listed <- price
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
