package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Version  string
	Revision string
)

var (
	configPath string
	c          Config
)

var rootCommand = &cobra.Command{
	Use: "bookstack",
}

func init() {
	cobra.OnInitialize(func() {
		if len(configPath) > 0 {
			viper.SetConfigFile(configPath)
		} else {
			viper.AddConfigPath(".")
			viper.SetConfigName("config")
		}
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		if err := viper.ReadInConfig(); err != nil {
			panic(err)
		}
		if err := viper.Unmarshal(&c); err != nil {
			panic(err)
		}
	})

	rootCommand.AddCommand(serveCommand())
}

func Execute() error {
	return rootCommand.Execute()
}
