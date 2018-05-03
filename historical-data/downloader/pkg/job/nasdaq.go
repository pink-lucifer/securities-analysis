package job

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/conf"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/parser"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"github.com/spf13/viper"
	"log"
)

var nasdaqfile string

func init() {

}

func InitNasdaq(config *conf.Config) {
	initDB(config)
	amexfile = viper.GetString("nasdaq.symfile")
}

func RumNasdaqSymbolParse() {
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

	csvReader, err := util.ReadCSV(amexfile)
	if err != nil {
		panic(err)
	}
	defer csvReader.Close()

	log.Print("Start processing csv file!")
	go util.ProcessCSV(csvReader, ch, eof)

	log.Print("Start parsing csv file record!")
	go parser.ParseAndWrapNasdaqSymbol(ch, eof, listed, done)

	go persistListedSymbol(sqlDb, listed, done, completed)
	<-completed
}
