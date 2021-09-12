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
	"time"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"

	"github.com/KushagraIndurkhya/fetch/core"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download <URL> <filename>",
	Short: "Download from url",
	Long: `
	Download a file from a download url
	Make sure you have at least twice free space 
	on the disk as the size of the file you are downloading
	`,
	Run: func(cmd *cobra.Command, args []string) {
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
	downloadCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fetch.yaml)")
	downloadCmd.PersistentFlags().String("path", "", "Specify Download Location of the file")
	downloadCmd.PersistentFlags().Bool("verbose", false, "Specify Verbosity of the output")
	downloadCmd.PersistentFlags().Int("threads", 20, "Specify Number of threads to be used")
	downloadCmd.PersistentFlags().Bool("seq", false, "Download the file sequentially instead of parallel downloading")
}
