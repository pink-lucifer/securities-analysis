package persist

import "testing"

func TestGetDb(t *testing.T) {
	cfg := &DataSourceConfig{
		DsType: "mysql",
		Url:    "security:security8888@tcp(127.0.0.1:3306)/fin_security?parseTime=true&allowNativePasswords=True",
	}

	CompleteConfig(cfg)

	sqlxDb := GetDb()
	defer sqlxDb.Close()

	_, err := sqlxDb.Exec("SELECT 1 FROM DUAL")
	if err != nil {
		t.Fatal(err)
	}
}
