package cmd

import (
	goflag "flag"
	"github.com/pink-lucifer/securities-analysis/historical-data/downloader/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"io"
	"strings"
)

var cfgFile string

func NewRootCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	rootCmd := cobra.Command{
		Use:   "downloader",
		Short: "downloader for security information",
		Long: `To get started run the server subcommand which will start a server
on localhost:10000:
    downloader server
Then you can hit it with the client:
    downloader client foo bar baz
Or over HTTP 1.1 with curl:
    curl -X POST -k https://localhost:10000/v1/echo -d '{"value": "foo"}'
`,
		Run: func(cmd *cobra.Command, args []string) {}}

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.
	//rootCmd.PersistentFlags().StringVarP(&cfgFile, "config.file", "f", "./config.yaml", "config file (default is /etc/notify/config.yaml)")

	rootCmd.AddCommand(ServerCmd)
	rootCmd.AddCommand(ClientCmd)
	rootCmd.AddCommand(JobCmd)
	rootCmd.AddCommand(VersionCmd)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	viper.BindPFlags(rootCmd.PersistentFlags())
	return &rootCmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName(util.DEFAULT_CONFIG_FILE_NAME) // name of config file (without extension)
		viper.AddConfigPath(util.DEFAULT_CONFIG_FILE_DIR)  // adding home directory as first search path
		viper.SetConfigType(util.DEFAULT_CONFIG_FILE_TYPE)
	}
	viper.AutomaticEnv() // read in environment variables that match
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	setupViper()
}

func setupViper() {
	viper.SetEnvPrefix(util.ENV_PREFXI)
	// Replaces '-' in flags with '_' in env variables
	// e.g. iso-url => $ENVPREFIX_ISO_URL
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
	//setFlagsUsingViper()
}
