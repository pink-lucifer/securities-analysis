package persist

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type TblListedSymbol struct {
	Id                    int64
	Uuid                  string
	Listed_market         uint16
	Symbol                string
	Name                  sql.NullString
	IpoYear               uint16
	Sector                sql.NullString
	Industry              sql.NullString
	SummaryQuote          sql.NullString
	Created_datetime      mysql.NullTime
	Last_updated_datetime mysql.NullTime
}
