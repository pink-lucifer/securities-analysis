package cmd

import (
	pkgcmd "github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/cmd"
	"os"
)

func Run() error {
	cmd := pkgcmd.NewRootCommand(os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}
