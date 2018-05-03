package util

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCSV(filePath string) (io.ReadCloser, error) {
	csvFile, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	return csvFile, nil
}

func ProcessCSV(rc io.Reader, ch chan<- []string, eof chan<- int) {
	r := csv.NewReader(rc)
	if _, err := r.Read(); err != nil { //read header
		log.Fatal(err)
	}
	var total int
	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			if err != csv.ErrFieldCount {
				log.Println(err)
				continue
			}
		}
		if len(rec) == 0 {
			continue
		}
		ch <- rec
		total++
	}
	eof <- total
}
