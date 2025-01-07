package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite 驅動
)

var DB *sql.DB

// InitDB 初始化資料庫連接
func InitDB(dbFile string) {
	var err error

	// 使用內存數據庫或檔案數據庫
	DB, err = sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 創建文章表
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	log.Println("Database initialized successfully!")
}
