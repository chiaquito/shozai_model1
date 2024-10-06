package main

import (
	// "github.com/labstack/echo/v4"

	"log"

	"github.com/go-playground/validator"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"shozai_model1/config"
	"shozai_model1/handler"
	"shozai_model1/repository"
	"shozai_model1/route"
	"shozai_model1/usecase"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {

	conf := config.New()

	db, err := gorm.Open(mysql.Open(conf.DB.FormatDSN()))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
		panic(err)
	}

	defer func() {
		sqlDB, _ := db.DB()
		_ = sqlDB.Close()
	}()

	r := repository.New(db)
	u := usecase.New(r)
	h := handler.NewHandler(u)

	e := route.NewRouter(h)

	e.Validator = &Validator{validator: validator.New()}

	e.Logger.Fatal(e.Start(":1323"))

}
