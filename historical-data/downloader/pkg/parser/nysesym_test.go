package parser

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"log"
	"testing"
)

func TestParseNyseSymbol(t *testing.T) {
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

	csvReader, err := util.ReadCSV(nyseCsvFile)
	if err != nil {
		panic(err)
	}
	defer csvReader.Close()

	log.Print("Start processing csv file!")
	go util.ProcessCSV(csvReader, ch, eof)

	log.Print("Start parsing csv file record!")
	go ParseNyseSymbol(ch, eof, listed, done)

	go ProcessNyseSymbol(listed, done, completed)
	<-completed
}

func TestParseAndWrapNyseSymbol(t *testing.T) {
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

	csvReader, err := util.ReadCSV(nyseCsvFile)
	if err != nil {
		panic(err)
	}
	defer csvReader.Close()

	log.Print("Start processing csv file!")
	go util.ProcessCSV(csvReader, ch, eof)

	log.Print("Start parsing csv file record!")
	go ParseAndWrapNyseSymbol(ch, eof, listed, done)

	go ProcessWrappedNyseSymbol(listed, done, completed)
	<-completed
}

func ProcessNyseSymbol(listed <-chan *model.ListedSymbol, eofCh <-chan int, done chan<- bool) {
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

func ProcessWrappedNyseSymbol(listed <-chan *persist.TblListedSymbol, eofCh <-chan int, done chan<- bool) {
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
