package main

import (
	"fmt"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/cmd"
	"os"
)

func main() {
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
