package config

import (
	"fmt"
	"time"

	env "github.com/caarlos0/env/v11"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type config struct {
	DB   *mysql.Config
	Gorm *gorm.Config
}

func New() *config {

	dbConf, err := newDBConf()
	if err != nil {
		fmt.Println(err)
	}
	gormConf := &gorm.Config{
		CreateBatchSize: 1000,
	}

	return &config{
		DB:   dbConf,
		Gorm: gormConf,
	}
}

type dbConf struct {
	Host     string `env:"DB_HOST" envDefault:"localhost"`
	Port     string `env:"DM_PORT" envDefault:"3306"`
	Name     string `env:"DB_NAME" envDefault:"db"`
	User     string `env:"DB_USER" envDefault:"root"`
	PassWord string `env:"DB_PASSWORD" envDefault:"password"`
}

func newDBConf() (*mysql.Config, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var dbCfg dbConf
	if err := env.Parse(&dbCfg); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &mysql.Config{
		DBName:               dbCfg.Name,
		User:                 dbCfg.User,
		Passwd:               dbCfg.PassWord,
		Addr:                 fmt.Sprintf("%s:%s", dbCfg.Host, dbCfg.Port),
		Net:                  "tcp",
		ParseTime:            true,
		Collation:            "utf8mb4_unicode_ci",
		Loc:                  jst,
		AllowNativePasswords: true,
	}, nil
}
