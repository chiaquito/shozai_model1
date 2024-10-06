package config

import (
	"time"

	"github.com/go-sql-driver/mysql"
)

type config struct {
	DB *mysql.Config
}

func New() *config {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// エラーハンドリング
	}
	dbConf := mysql.Config{
		DBName:    "db",
		User:      "root",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
		Params:    map[string]string{
			// "usessl": "false",
			// "allowPublicKeyRetrieval": "true",
		},
	}

	return &config{
		DB: &dbConf,
	}
}
