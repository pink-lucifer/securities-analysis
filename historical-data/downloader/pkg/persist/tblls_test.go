package persist

import (
	"database/sql"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"testing"
)

func TestInsertListedSymbol(t *testing.T) {
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

	symbol := &TblListedSymbol{
		Uuid:              uid,
		Listed_market:     uint16(9999),
		Symbol:            "TSTSYM",
		Name:              sql.NullString{String: "Test Name", Valid: true},
		Sector:            sql.NullString{String: "Test Sector", Valid: true},
		Industry:          sql.NullString{String: "Test Industry", Valid: true},
		Summary_quote_url: sql.NullString{String: "Test Summary_quote_url", Valid: true},
		Ipo_year:          uint16(1999),
	}
	tx, err := sqlxDb.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	id, ra, err := InsertListedSymbol(tx, symbol)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("InsertListedSymbol done successfully, id: %d, row affected: %d", id, ra)

	tx.Rollback()
}

func TestSelectListedSymbolByMktAndSym(t *testing.T) {

	cfg := &DataSourceConfig{
		DsType: "mysql",
		Url:    "security:security8888@tcp(127.0.0.1:3306)/fin_security?parseTime=true&allowNativePasswords=True",
	}

	CompleteConfig(cfg)

	sqlxDb := GetDb()
	defer sqlxDb.Close()

	mkt := uint16(9999)
	sym := "TEST"
	symbols, err := SelectListedSymbolByMktAndSym(sqlxDb, mkt, sym)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(symbols)
}
