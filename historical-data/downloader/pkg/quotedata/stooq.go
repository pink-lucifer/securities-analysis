package quotedata

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"time"
)

const (
	baseUrl        = "https://stooq.com/q/d/l/?s="
	suffix         = "&i=d"
	midfix         = ".US"
	amexDir        = "../../data/listed-mkt/amex/historical/"
	nasdaqDir      = "../../data/listed-mkt/nasdaq/historical/"
	nyseDir        = "../../data/listed-mkt/nyse/historical/"
	fileTypeSuffix = ".csv"
)

func HistoricalQuoteInitialDownload(mkt model.Market, symbol string, now time.Time) error {
	url := baseUrl + symbol + midfix + suffix
	var dir string
	switch mkt {
	case model.NYSE:
		dir = nyseDir
		break
	case model.AMEX:
		dir = amexDir
		break
	case model.NASDAQ:
		dir = nasdaqDir
		break
	}
	nowStr := now.Format("2006-01-02-15-04-05")
	fileName := dir + symbol + midfix + "." + nowStr + fileTypeSuffix
	err := util.DownloadToFile(url, fileName)
	return err
}
