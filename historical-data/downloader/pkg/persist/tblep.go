package persist

import (
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	TBL_EOD_PRICE_INSERT            = "insert into eod_price(uuid, listed_market, symbol, trade_date, open_price, close_price, highest_price, lowest_price, volume) values (?,?,?,?,?,?,?,?,?)"
	TBL_EOD_PRICE_SELECT_BY_MKT_SYM = "select id, uuid, listed_market, symbol, trade_date, open_price, close_price, highest_price, lowest_price, volume, created_timestamp, last_updated_timestamp" +
		" from eod_price where listed_market =:market and symbol =:symbol"
	TBL_EOD_PRICE_SELECT_BY_DATE_MKT_SYM = "select id, uuid, listed_market, symbol, name, ipo_year, sector, industry, summary_quote_url, created_timestamp, last_updated_timestamp" +
		" from eod_price where trade_date:=date and listed_market =:market and symbol =:symbol"
)

func InsertEodPrice(tx *sqlx.Tx, price *TblEodPrice) (int64, int64, error) {
	r, err := tx.Exec(TBL_EOD_PRICE_INSERT, price.Uuid, price.Listed_market, price.Symbol, price.Trade_date, price.Open_price, price.Close_price, price.Highest_price, price.Lowest_price, price.Volume)
	if err != nil {
		return 0, 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	ra, err := r.RowsAffected()
	if err != nil {
		return id, 0, err
	}
	return id, ra, nil
}

func SelectEodPriceByDateAndMktAndSym(db *sqlx.DB, trdDate time.Time, mkt uint16, symbol string) ([]TblEodPrice, error) {
	p := map[string]interface{}{"date": trdDate, "market": mkt, "symbol": symbol}
	prices := []TblEodPrice{}
	stmt, err := db.PrepareNamed(TBL_EOD_PRICE_SELECT_BY_DATE_MKT_SYM)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&prices, p)
	if err != nil {
		return nil, err
	}

	return prices, nil
}

func SelectEodPriceByMktAndSym(db *sqlx.DB, mkt uint16, symbol string) ([]TblEodPrice, error) {
	p := map[string]interface{}{"market": mkt, "symbol": symbol}
	prices := []TblEodPrice{}
	stmt, err := db.PrepareNamed(TBL_EOD_PRICE_SELECT_BY_MKT_SYM)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&prices, p)
	if err != nil {
		return nil, err
	}

	return prices, nil
}
