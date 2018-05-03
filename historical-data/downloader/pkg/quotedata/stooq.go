package quotedata

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"github.com/spf13/viper"
	"strings"
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
		dir = viper.GetString("nyse.dir")
		if dir == "" || strings.TrimSpace(dir) == "" {
			dir = nyseDir
		}
		break
	case model.AMEX:
		dir = viper.GetString("amex.dir")
		if dir == "" || strings.TrimSpace(dir) == "" {
			dir = amexDir
		}
		break
	case model.NASDAQ:
		dir = viper.GetString("nasdaq.dir")
		if dir == "" || strings.TrimSpace(dir) == "" {
			dir = nasdaqDir
		}
		break
	}
	err := util.EnsureDirExist(dir)
	if err != nil {
		return err
	}
	nowStr := now.Format("2006-01-02-15-04-05")
	fileName := viper.GetString("filename")
	if fileName == "" || strings.TrimSpace(fileName) == "" {
		fileName = dir + symbol + midfix + "." + nowStr + fileTypeSuffix
	}

	err = util.DownloadToFile(url, fileName)
	return err
}
