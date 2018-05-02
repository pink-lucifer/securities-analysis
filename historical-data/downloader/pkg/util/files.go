package util

import (
	"io/ioutil"
	"log"
	"net/http"
)

func HttpDownload(uri string) ([]byte, error) {
	log.Printf("HttpDownload From: %s.\n", uri)
	res, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	d, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ReadFile: Size of download: %d\n", len(d))
	return d, err
}

func WriteFile(dst string, d []byte) error {
	log.Printf("WriteFile: Size of download: %d\n", len(d))
	err := ioutil.WriteFile(dst, d, 0444)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func DownloadToFile(uri string, dst string) error {
	log.Printf("DownloadToFile From: %s.\n", uri)
	data, err := HttpDownload(uri)
	if err != nil {
		return err
	}
	log.Printf("downloaded %s.\n", uri)

	err = WriteFile(dst, data)
	if err != nil {
		return err
	}
	log.Printf("saved %s as %s\n", uri, dst)

	return nil
}
