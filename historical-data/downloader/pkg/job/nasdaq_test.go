package job

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/conf"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"github.com/spf13/viper"
	"testing"
)

func TestRunNasdaqSymbolParse(t *testing.T) {
	viper.Set("nasdaq.symfile", "../../data/listed-mkt/nasdaq/companylist-nasdaq.csv")
	cfg := &conf.Config{
		DataSourceConfig: &persist.DataSourceConfig{
			DsType:  "mysql",
			Url:     "security:security8888@tcp(127.0.0.1:3306)/fin_security?parseTime=true&allowNativePasswords=True",
			MaxIdle: 1,
			MaxOpen: 160,
		},
	}

	InitNasdaq(cfg)

	defer sqlDb.Close()
	RunNyseSymbolParse()
}
