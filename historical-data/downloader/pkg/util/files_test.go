package util

import (
	"os"
	"testing"
)

func TestDownloadToFile(t *testing.T) {
	url := "https://stooq.com/q/d/l/?s=AAPL.US&i=d"
	dst := "../../data/TestDownloadToFile.csv"
	err := DownloadToFile(url, dst)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Remove(dst)
	if err != nil {
		t.Fatal(err)
	}
}

func TestHttpDownload(t *testing.T) {
	url := "https://stooq.com/q/d/l/?s=AAPL.US&i=d"
	data, err := HttpDownload(url)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Downloaded data: %s", string(data))
}

func TestWriteFile(t *testing.T) {
	dst := "../../data/TestWriteFile.csv"
	data := []byte("1,2,3,4,5,6,7,8")
	err := WriteFile(dst, data)
	if err != nil {
		t.Fatal(err)
	}

	err = os.Remove(dst)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnsureDirExist(t *testing.T) {
	path := "../../data/listed-mkt/nasdaq/historical/"
	err := EnsureDirExist(path)
	if err != nil {
		t.Fatal(err)
	}
}
