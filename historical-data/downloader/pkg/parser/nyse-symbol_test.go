package parser

import (
	"../model"
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

	csvReader, err := ReadCSV(nyseCsvFile)
	if err != nil {
		panic(err)
	}
	defer csvReader.Close()

	log.Print("Start processing csv file!")
	go ProcessCSV(csvReader, ch, eof)

	log.Print("Start parsing csv file record!")
	go ParseNyseSymbol(ch, eof, listed, done)

	go ProcessNyseSymbol(listed, done, completed)
	<-completed
}


func ProcessNyseSymbol(listed <-chan *model.ListedSymbol, eofCh <-chan int, done chan<-bool)  {
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