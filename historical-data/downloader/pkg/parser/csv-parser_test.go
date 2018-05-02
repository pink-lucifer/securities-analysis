package parser

import "testing"

func TestReadCSV(t *testing.T) {
	csvFile, err:=ReadCSV(amexCsvFile)
	defer csvFile.Close()
	if err != nil{
		t.Error(err)
	}
}

func TestProcessCSV(t *testing.T) {

}
