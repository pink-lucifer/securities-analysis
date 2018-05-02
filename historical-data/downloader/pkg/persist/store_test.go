package persist

import (
	"../util"
	"database/sql"
	"testing"
)

func TestInsert(t *testing.T) {
	uid, err := util.GenerateUUIDV4()
	if err != nil {
		t.Fail()
	}
	cfg := &DataSourceConfig{
		DsType: "mysql",
		Url:    "security:security8888@tcp(127.0.0.1:3306)/fin_security?parseTime=true",
	}

	CompleteConfig(cfg)

	sqlxDb := GetDb()

	symbol := &TblListedSymbol{
		Uuid:          uid,
		Listed_market: uint16(99),
		Symbol:        "TEST",
		Name:          sql.NullString{String: "Test Name", Valid: true},
		Sector:        sql.NullString{String: "Test Sector", Valid: true},
		Industry:      sql.NullString{String: "Test Industry", Valid: true},
		SummaryQuote:  sql.NullString{String: "Test SummaryQuote", Valid: true},
		IpoYear:       uint16(1999),
	}
	tx, err := sqlxDb.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	id, ra, err := Insert(tx, symbol)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Insert done successfully, id: %d, row affected: %d", id, ra)

	tx.Rollback()
}
