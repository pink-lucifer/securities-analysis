package job

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"log"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/conf"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/parser"
)

func init() {

}

func InitStooq(config *conf.Config) {
	initDB(config)
}

func RunStooqEodQuote(market model.Market, sym string, fileName string) {
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

	csvReader, err := util.ReadCSV(fileName)
	if err != nil {
		panic(err)
	}
	defer csvReader.Close()

	log.Print("Start processing stooq csv file!")
	go util.ProcessCSV(csvReader, ch, eof)

	log.Print("Start parsing stooq csv file record!")
	go parser.ParseAndWrapEodPrice(sym, market, ch, eof, listed, done)

	go pooledPersistEodPrice(sqlDb, listed, done, completed)
	<-completed
}
