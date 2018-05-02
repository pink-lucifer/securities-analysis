package persist

import "github.com/jmoiron/sqlx"

const (
	TBL_LISTED_SYMBOL_INSERT            = "insert into listed_symbol(uuid, listed_market, symbol, name, ipo_year, sector, industry, summary_quote_url) values (?,?,?,?,?,?,?,?)"
	TBL_LISTED_SYMBOL_SELECT_BY_MKT_SYM = "select id, uuid, listed_market, symbol, name, ipo_year, sector, industry, summary_quote_url, created_timestamp, last_updated_timestamp" +
		" from listed_symbol where listed_market =:market and symbol =:symbol"
)

func Insert(tx *sqlx.Tx, symbol *TblListedSymbol) (int64, int64, error) {
	r, err := tx.Exec(TBL_LISTED_SYMBOL_INSERT, symbol.Uuid, symbol.Listed_market, symbol.Symbol, symbol.Name, symbol.Ipo_year, symbol.Sector, symbol.Industry, symbol.Summary_quote_url)
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

func SelectByMktAndSym(db *sqlx.DB, mkt uint16, symbol string) ([]TblListedSymbol, error) {
	p := map[string]interface{}{"market": mkt, "symbol": symbol}
	symbols := []TblListedSymbol{}
	stmt, err := db.PrepareNamed(TBL_LISTED_SYMBOL_SELECT_BY_MKT_SYM)
	if err != nil {
		return nil, err
	}

	err = stmt.Select(&symbols, p)
	if err != nil {
		return nil, err
	}

	return symbols, nil
}
