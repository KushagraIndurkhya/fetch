package core

import (
	"fmt"
	"os"
)

func Cleanup() {
	CreateSqliteDb()
	db := OpenSqliteDb()
	defer db.Close()
	list := GetHistoryList(db, 10)
	for _, v := range list {
		fmt.Println(fmt.Sprint(v.Id) + " Deleting " + v.FilePath)
		if err := os.Remove(v.FilePath); err != nil {
			fmt.Println("Error:" + fmt.Sprint(err) + "in deleting " + v.FilePath)
		}
		DeleteHistory(db, v.Id)
	}

}
