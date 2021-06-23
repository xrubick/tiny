package cmd

import (
	"fmt"
	"os"
	"time"
	"context"

	"github.com/vultr/govultr/v2"
	"golang.org/x/oauth2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "tiny",
	Short: "Tiny is just a cli tool for the vultr api",
	Long: ``,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tiny.yaml)")
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(instanceCmd)


}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".tiny" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tiny")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
