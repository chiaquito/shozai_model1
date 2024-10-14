package config

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	DB   *mysql.Config
	Gorm *gorm.Config
}

func New() *config {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		// エラーハンドリング
	}
	dbConf := &mysql.Config{
		DBName: "db",
		User:   "root",
		Passwd: "password",
		// Addr:      "localhost:3306",
		Addr:      "db:3306",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
		Params:    map[string]string{
			// "usessl": "false",
			// "allowPublicKeyRetrieval": "true",
		},
	}

	gormConf := &gorm.Config{
		CreateBatchSize: 1000,
	}

	return &config{
		DB:   dbConf,
		Gorm: gormConf,
	}
}
