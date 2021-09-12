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

	"github.com/KushagraIndurkhya/fetch/core"
	"github.com/spf13/cobra"
)

// historyCmd represents the history command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Fetch your download history",
	Long:  `Get details of the files downloaded using fetch`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Fetch Download History as of " + fmt.Sprint(time.Now().Format("2006-01-02 15:04:05")) + "\n")
		core.CreateSqliteDb()
		db := core.OpenSqliteDb()
		no_of_results, err := cmd.Flags().GetInt("list")
		if err != nil {
			fmt.Println(err)
		}
		list := core.GetHistoryList(db, no_of_results)
		for _, v := range list {
			fmt.Printf("%d\t%s\t%s\t%s\t%s\n", v.Id, v.FileName, v.FileSize, v.Date, v.Success)
		}

	},
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clear your downloading history",
	Long:  `Delete all records from your history sqlite db`,
	Run: func(cmd *cobra.Command, args []string) {

		core.CreateSqliteDb()
		db := core.OpenSqliteDb()
		err := core.DeleteAllHistory(db)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("History cleared successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(historyCmd)
	historyCmd.AddCommand(cleanCmd)
	historyCmd.PersistentFlags().Int("list", 10, "Specify Number of Rows in result")
}
