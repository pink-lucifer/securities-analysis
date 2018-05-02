package persist

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type TblListedSymbol struct {
	Id                     int64
	Uuid                   string
	Listed_market          uint16
	Symbol                 string
	Name                   sql.NullString
	Ipo_year               uint16
	Sector                 sql.NullString
	Industry               sql.NullString
	Summary_quote_url      sql.NullString
	Created_timestamp      mysql.NullTime
	Last_updated_timestamp mysql.NullTime
}
