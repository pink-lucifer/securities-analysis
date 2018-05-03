package persist

import (
	"database/sql"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"testing"
	"time"
)

func TestInsertEodPrice(t *testing.T) {
	uid, err := util.GenerateUUIDV4()
	if err != nil {
		t.Fail()
	}
	cfg := &DataSourceConfig{
		DsType: "mysql",
		Url:    "security:security8888@tcp(127.0.0.1:3306)/fin_security?parseTime=true&allowNativePasswords=True",
	}

	CompleteConfig(cfg)

	sqlxDb := GetDb()
	defer sqlxDb.Close()

	price := &TblEodPrice{
		Uuid:          uid,
		Listed_market: uint16(9999),
		Symbol:        "TSTSYM",
		Trade_date:    time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local),
		Open_price:    sql.NullFloat64{Float64: float64(10.0), Valid: true},
		Close_price:   sql.NullFloat64{Float64: float64(10.98), Valid: true},
		Highest_price: sql.NullFloat64{Float64: float64(11.12), Valid: true},
		Lowest_price:  sql.NullFloat64{Float64: float64(9.75), Valid: true},
		Volume:        sql.NullFloat64{Float64: float64(10000000.0), Valid: true},
	}
	tx, err := sqlxDb.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	id, ra, err := InsertEodPrice(tx, price)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("InsertEodPrice done successfully, id: %d, row affected: %d", id, ra)

	tx.Rollback()
}

func TestSelectEodPriceByMktAndSym(t *testing.T) {

	cfg := &DataSourceConfig{
		DsType: "mysql",
		Url:    "security:security8888@tcp(127.0.0.1:3306)/fin_security?parseTime=true&allowNativePasswords=True",
	}

	CompleteConfig(cfg)

	sqlxDb := GetDb()
	defer sqlxDb.Close()
	mkt := uint16(9999)
	sym := "TEST"
	prices, err := SelectEodPriceByMktAndSym(sqlxDb, mkt, sym)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(prices)
}

func TestSelectEodPriceByDateAndMktAndSym(t *testing.T) {

	cfg := &DataSourceConfig{
		DsType: "mysql",
		Url:    "security:security8888@tcp(127.0.0.1:3306)/fin_security?parseTime=true&allowNativePasswords=True",
	}

	CompleteConfig(cfg)

	sqlxDb := GetDb()
	defer sqlxDb.Close()

	mkt := uint16(9999)
	sym := "TEST"
	date := time.Date(2018, 1, 1, 0, 0, 0, 0, time.Local)
	prices, err := SelectEodPriceByDateAndMktAndSym(sqlxDb, date, mkt, sym)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(prices)
}
