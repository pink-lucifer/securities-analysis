package parser

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"log"
	"testing"
)

func TestParseAmexSymbol(t *testing.T) {
	ch := make(chan []string)
	defer close(ch)
	eof := make(chan int, 1)
	defer close(eof)
	listed := make(chan *model.ListedSymbol)
	defer close(listed)
	done := make(chan int, 1)
	defer close(done)
	completed := make(chan bool, 1)
	defer close(completed)

	csvReader, err := util.ReadCSV(AmexCsvFile)
	if err != nil {
		panic(err)
	}
	defer csvReader.Close()

	log.Print("Start processing csv file!")
	go util.ProcessCSV(csvReader, ch, eof)

	log.Print("Start parsing csv file record!")
	go ParseAmexSymbol(ch, eof, listed, done)

	go ProcessAmexSymbol(listed, done, completed)
	<-completed
}

func TestParseAndWrapAmexSymbol(t *testing.T) {
	ch := make(chan []string)
	defer close(ch)
	eof := make(chan int, 1)
	defer close(eof)
	listed := make(chan *persist.TblListedSymbol)
	defer close(listed)
	done := make(chan int, 1)
	defer close(done)
	completed := make(chan bool, 1)
	defer close(completed)

	csvReader, err := util.ReadCSV(AmexCsvFile)
	if err != nil {
		panic(err)
	}
	defer csvReader.Close()

	log.Print("Start processing csv file!")
	go util.ProcessCSV(csvReader, ch, eof)

	log.Print("Start parsing csv file record!")
	go ParseAndWrapAmexSymbol(ch, eof, listed, done)

	go ProcessWrappedAmexSymbol(listed, done, completed)
	<-completed
}

func ProcessAmexSymbol(listed <-chan *model.ListedSymbol, eofCh <-chan int, done chan<- bool) {
	var total, processed int
	for {
		select {
		case symbol := <-listed:
			log.Print(symbol)
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

func ProcessWrappedAmexSymbol(listed <-chan *persist.TblListedSymbol, eofCh <-chan int, done chan<- bool) {
	var total, processed int
	for {
		select {
		case symbol := <-listed:
			log.Print(symbol)
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
