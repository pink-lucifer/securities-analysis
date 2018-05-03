package job

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/conf"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"github.com/spf13/viper"
	"testing"
)

func TestRunNyseSymbolParse(t *testing.T) {
	viper.Set("nyse.symfile", "../../data/listed-mkt/nyse/companylist-nyse.csv")
	cfg := &conf.Config{
		DataSourceConfig: &persist.DataSourceConfig{
			DsType:  "mysql",
			Url:     "security:security8888@tcp(127.0.0.1:3306)/fin_security?parseTime=true&allowNativePasswords=True",
			MaxIdle: 1,
			MaxOpen: 160,
		},
	}

	InitNyse(cfg)

	defer sqlDb.Close()
	RunNyseSymbolParse()
}
