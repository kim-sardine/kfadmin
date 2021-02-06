package cmd

import (
	"fmt"
	"os"

	"github.com/kim-sardine/kfadmin/client"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var c *client.KfClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kfadmin",
	Short: "CLI Tool for Kubeflow admin",
	Long: `kfadmin is a CLI tool for kubeflow admin.
Please check your k8s, kubeflow versions before use it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kfadmin.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	c = &client.KfClient{}
	c.LoadClientset()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".kfadmin" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".kfadmin")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
