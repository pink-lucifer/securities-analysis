package cmd

import "github.com/spf13/cobra"

// ServerCmd represents the server command
var (
	ServerCmd = &cobra.Command{
		Use:   "server",
		Short: "Launches the web server on https://localhost:10000",
		Run: func(cmd *cobra.Command, args []string) {
			server()
		},
	}
)

func server() {

}
