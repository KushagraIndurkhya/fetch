package core

import (
	"database/sql"
	"log"

	"os"
	// "time"

	_ "github.com/mattn/go-sqlite3"
)

var DB_NAME string = "/history.db"

func get_DB_path() string {
	if os.Getenv("HOME") != "" {
		return os.Getenv("HOME") + "/.config/fetch"
	} else {
		return "/tmp"
	}
}

type History struct {
	Id       int
	Date     string
	Success  string
	FilePath string
	FileName string
	FileSize string
}

func OpenSqliteDb() *sql.DB {
	// fmt.Println("Opening SQLite database...")
	db, err := sql.Open("sqlite3", get_DB_path()+DB_NAME)
	if err != nil {
		log.Println("Error opening SQLite database: ", err)
	}

	return db
}

func CreateSqliteDb() {
	// fmt.Println("Creating SQLite database...")
	os.MkdirAll(get_DB_path(), 0755)
	if _, err := os.Stat(get_DB_path() + DB_NAME); os.IsNotExist(err) {
		os.Create(get_DB_path() + DB_NAME)
	}
	history_db, err := sql.Open("sqlite3", get_DB_path()+DB_NAME)
	if err != nil {
		panic(err)
	}
	defer history_db.Close()
	Initialize_table(history_db)

}

func Initialize_table(db *sql.DB) {

	createHistoryTable := `CREATE TABLE IF NOT EXISTS history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date TEXT,
		Success TEXT,
		FilePath TEXT,
		FileName TEXT,
		FileSize TEXT
		);`
	// log.Println("Initializing history DB...")
	_, err := db.Exec(createHistoryTable)
	if err != nil {
		log.Println("Error creating history table: ", err)
	}
	// log.Println("History table initialized.")

}

func InsertHistory(db *sql.DB, hist History) {
	insertHistoryQuery := `INSERT INTO history (date, Success,FilePath,FileName,FileSize) VALUES (?, ?, ?, ?, ?)`
	_, err := db.Exec(insertHistoryQuery, hist.Date, hist.Success, hist.FilePath, hist.FileName, hist.FileSize)
	if err != nil {
		log.Println("Error inserting history: ", err)
	}
}

func DeleteHistory(db *sql.DB, id int) {
	deleteHistoryQuery := `DELETE FROM history WHERE Id=?`
	_, err := db.Exec(deleteHistoryQuery, id)
	if err != nil {
		log.Println("Error deleting history: ", err)
	}
}

func DeleteAllHistory(db *sql.DB) error {
	deleteAllHistoryQuery := `DELETE FROM history`
	_, err := db.Exec(deleteAllHistoryQuery)
	if err != nil {
		log.Println("Error deleting all history: ", err)
	}
	return err
}

func GetHistoryList(db *sql.DB, limit int) []History {
	getHistoryQuery := `SELECT * FROM history ORDER BY date`
	rows, err := db.Query(getHistoryQuery)
	if err != nil {
		log.Println("Error getting history: ", err)
	}
	defer rows.Close()
	var history []History
	i := 0
	for rows.Next() {
		if i >= limit {
			break
		}
		var h History
		err := rows.Scan(&h.Id, &h.Date, &h.Success, &h.FilePath, &h.FileName, &h.FileSize)
		if err != nil {
			log.Println("Error getting history: ", err)
		}
		history = append(history, h)
		i++
	}
	return history
}

// func checkErr(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
