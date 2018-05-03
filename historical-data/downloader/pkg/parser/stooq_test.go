package parser

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"log"
	"testing"
)

const aaplEodFile = "../../data/test/AAPL.US.2018-05-01.csv"

func TestParseAndWrapEodPrice(t *testing.T) {
	sym := "AAPL"
	ch := make(chan []string)
	defer close(ch)
	eof := make(chan int, 1)
	defer close(eof)
	listed := make(chan *persist.TblEodPrice)
	defer close(listed)
	done := make(chan int, 1)
	defer close(done)
	completed := make(chan bool, 1)
	defer close(completed)

	csvReader, err := util.ReadCSV(aaplEodFile)
	if err != nil {
		panic(err)
	}
	defer csvReader.Close()

	log.Print("Start processing csv file!")
	go util.ProcessCSV(csvReader, ch, eof)

	log.Print("Start parsing csv file record!")
	go ParseAndWrapEodPrice(sym, model.NASDAQ, ch, eof, listed, done)

	go ProcessWrappedEodPrice(listed, done, completed)
	<-completed
}

func ProcessWrappedEodPrice(listed <-chan *persist.TblEodPrice, eofCh <-chan int, done chan<- bool) {
	var total, processed int
	for {
		select {
		case price := <-listed:
			log.Print(price)
			processed++
		case eof := <-eofCh:
			total = eof
		default:
			if total != 0 && processed != 0 && total == processed {
				done <- true
				return // differences between break and return
			}
		}
	}
}
