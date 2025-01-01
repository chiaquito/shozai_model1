package main

import (
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"shozai_model1/internal/config"
	"shozai_model1/internal/handler"
	"shozai_model1/internal/repository"
	"shozai_model1/internal/route"
	"shozai_model1/internal/usecase"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func main() {

	conf := config.New()

	db, err := gorm.Open(mysql.Open(conf.DB.FormatDSN()), conf.Gorm)
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323"},
	}))

	e.Logger.Fatal(e.Start(":1323"))

}
