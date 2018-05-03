package parser

import (
	"database/sql"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/model"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/persist"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"strconv"
	"strings"
	"time"
)

func WrapEodPrice(sym string, mkt model.Market, line []string) (*persist.TblEodPrice, error) {
	uid, err := util.GenerateUUIDV4()
	if err != nil {
		return nil, err
	}

	date, err := time.Parse("2006-01-02", line[0])
	if err != nil {
		return nil, err
	}

	open, _ := strconv.ParseFloat(line[1], 64)
	close, _ := strconv.ParseFloat(line[2], 64)
	high, _ := strconv.ParseFloat(line[3], 64)
	low, _ := strconv.ParseFloat(line[4], 64)
	var volume float64
	if len(line) > 5 {
		volume, _ = strconv.ParseFloat(line[5], 64)
	}

	return &persist.TblEodPrice{
		Uuid:          uid,
		Listed_market: uint16(mkt),
		Symbol:        sym,
		Trade_date:    date,
		Open_price:    sql.NullFloat64{Float64: open, Valid: true},
		Close_price:   sql.NullFloat64{Float64: close, Valid: true},
		Highest_price: sql.NullFloat64{Float64: high, Valid: true},
		Lowest_price:  sql.NullFloat64{Float64: low, Valid: true},
		Volume:        sql.NullFloat64{Float64: volume, Valid: true},
	}, nil
}

func WrapListedSymbol(mkt model.Market, line []string) (*persist.TblListedSymbol, error) {
	uid, err := util.GenerateUUIDV4()
	if err != nil {
		return nil, err
	}
	year, _ := strconv.Atoi(strings.TrimSpace(line[5]))

	return &persist.TblListedSymbol{
		Uuid:              uid,
		Listed_market:     uint16(mkt),
		Symbol:            strings.TrimSpace(line[0]),
		Name:              sql.NullString{String: strings.TrimSpace(line[1]), Valid: true},
		Ipo_year:          uint16(year),
		Sector:            sql.NullString{String: strings.TrimSpace(line[6]), Valid: true},
		Industry:          sql.NullString{String: strings.TrimSpace(line[7]), Valid: true},
		Summary_quote_url: sql.NullString{String: strings.TrimSpace(line[8]), Valid: true},
	}, nil
}
