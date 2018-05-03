package conf

import (
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
)

type Config struct {
	DataSourceConfig *persist.DataSourceConfig
}
