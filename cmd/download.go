/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		core.CreateSqliteDb()
		db := core.OpenSqliteDb()
		var hist *core.History = new(core.History)

		location, err := cmd.Flags().GetString("path")
		if err != nil || location == "" {
			location = viper.GetString("Default_Path")
		}
		chunks, err := cmd.Flags().GetInt("threads")
		cobra.CheckErr(err)
		verbose, err := cmd.Flags().GetBool("verbose")
		cobra.CheckErr(err)

		info := core.Make_info(
			args[0],
			args[1],
			location, chunks)
		seq, err := cmd.Flags().GetBool("seq")
		cobra.CheckErr(err)
		start_time := time.Now()

		fmt.Print(info)

		hist.FilePath = info.Location + "/" + info.NAME
		hist.FileName = info.NAME
		hist.FileSize = fmt.Sprint(info.Length)
		hist.Date = start_time.Format("2006-01-02 15:04:05")
		hist.Success = "Downloading"

		fmt.Println("Downloading Started")
		if seq {
			if core.Download_Seq(info, verbose) == nil {
				fmt.Printf("File %s downloaded in %f seconds ", args[1], time.Since(start_time).Seconds())
				hist.Success = "Success"
			} else {
				fmt.Printf("Something Went Wrong %s ", err)
				hist.Success = "Failed" + err.Error()
			}
		} else {
			if core.Download(info, verbose) == nil {
				fmt.Printf("File %s downloaded in %f seconds ", args[1], time.Since(start_time).Seconds())
				hist.Success = "Success"
			} else {
				fmt.Printf("Something Went Wrong %s ", err)
				hist.Success = "Failed" + err.Error()
			}
		}

		core.InsertHistory(db, *hist)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fetch.yaml)")
	rootCmd.PersistentFlags().String("path", "", "Specify Download Location of the file")
	rootCmd.PersistentFlags().Bool("verbose", false, "Specify Verbosity of the output")
	rootCmd.PersistentFlags().Int("threads", 20, "Specify Number of threads to be used")
	rootCmd.PersistentFlags().Bool("seq", false, "Download the file sequentially instead of parallel downloading")
}
