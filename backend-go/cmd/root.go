package cmd

import (
	"fmt"
	"os"

	"github.com/Hotpot-protocol1/hotpot-global/config"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var cfg config.Conf

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hotpot-backend",
	Short: "hotpot-backend will start hotpot backend service",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hotpot-backend.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".hotpot-backend" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".hotpot-backend")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Look for env vars with HOTPOT_ prefix
	viper.SetEnvPrefix("HOTPOT")
	viper.AllowEmptyEnv(true)

	// get all the keys from config.Conf object and then bind those keys to env vars
	envKeysMap := &map[string]interface{}{}
	if err := mapstructure.Decode(config.Conf{}, &envKeysMap); err != nil {
		fmt.Fprintln(os.Stderr, "Unable to decode config: ", err)
	}

	for k := range *envKeysMap {
		if bindErr := viper.BindEnv(k); bindErr != nil {
			fmt.Fprintln(os.Stderr, "Unable to bind env:", k)
		}
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to unmarshal configuration! ENV keys: %v, error %v", viper.AllKeys(), err)
		return
	}
}
