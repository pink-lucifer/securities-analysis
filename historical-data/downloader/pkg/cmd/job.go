package cmd

import (
	"github.com/spf13/cobra"
)

// ServerCmd represents the server command
var (
	JobCmd = &cobra.Command{
		Use:   "job",
		Short: "Launches the webserver on https://localhost:10000",
		Run: func(cmd *cobra.Command, args []string) {
			job()
		},
	}
)

func job() {

}
