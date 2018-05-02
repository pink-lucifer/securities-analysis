package cmd

import (
	"fmt"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/version"
	"github.com/spf13/cobra"
)

// ServerCmd represents the server command
var (
	VersionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the semantic version information.",
		Run: func(cmd *cobra.Command, args []string) {
			semanticVersion()
		},
	}
)

func semanticVersion() {
	fmt.Println(version.Version)
}
