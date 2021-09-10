/*
Copyright © 2021 Kushagra Indurkhya <kushagraindurkhya7@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"github.com/KushagraIndurkhya/fetch/core"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fetch <URL> <filename>",
	Short: "A CLI for fast downloads ",
	Long: `Fetch is a CLI for downloading files written in Go that enables user to get fast downloads by 
utilizing multiple threads and downloading file chunks in parallel`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		location, err := cmd.Flags().GetString("path")
		if err != nil || location == "" {
			location = viper.GetString("Default_Path")
		}
		chunks, err := cmd.Flags().GetInt("threads")
		cobra.CheckErr(err)
		verbose, err := cmd.Flags().GetBool("verbose")
		cobra.CheckErr(err)

		start_time := time.Now()
		info := core.Make_info(
			args[0],
			args[1],
			location, chunks)

		if core.Download(info, verbose) == nil {
			fmt.Printf("File %s downloaded in %f seconds ", args[1], time.Since(start_time).Seconds())
		} else {
			fmt.Printf("Something Went Wrong %s ", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fetch.yaml)")
	rootCmd.PersistentFlags().String("path", "", "Specify Download Location of the file")
	rootCmd.PersistentFlags().Bool("verbose", false, "Specify Verbosity of the output")
	rootCmd.PersistentFlags().Int("threads", 20, "Specify Number of threads to be used")
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

		// Search config in home directory with name ".fetch" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".fetch")

		viper.SetDefault("Default_Path", home+"/Downloads")

	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
