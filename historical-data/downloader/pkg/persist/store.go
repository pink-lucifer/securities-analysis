package persist

import "github.com/jmoiron/sqlx"

const TBL_INSERT_LISTED_SYMBOL = "insert into listed_symbol(uuid, listed_market, symbol, name, ipo_year, sector, industry, summary_quote_url) values (?,?,?,?,?,?,?,?)"

func Insert(tx *sqlx.Tx, symbol *TblListedSymbol) (int64, int64, error)  {
	r, err := tx.Exec(TBL_INSERT_LISTED_SYMBOL, symbol.Uuid, symbol.Listed_market, symbol.Symbol, symbol.Name, symbol.IpoYear, symbol.Sector, symbol.Industry, symbol.SummaryQuote)
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

