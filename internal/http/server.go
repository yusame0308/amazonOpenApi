package api

import (
	"amazonOpenApi/internal/http/gen"
	"amazonOpenApi/internal/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Run() {
	e := echo.New()
	// リクエストIDの設定
	e.Use(middleware.RequestID())
	// loggerの設定
	e.Use(middleware.Logger())
	// recoverの設定
	e.Use(middleware.Recover())
	// mysql connection
	dsn := "docker:docker@tcp(127.0.0.1:3306)/amazonOpenApi?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// Migrate the schema
	if err := db.AutoMigrate(&repository.AmazonData{}); err != nil {
		panic(err.Error())
	}
	gen.RegisterHandlers(e, NewApi(db))
	e.Logger.Fatal(e.Start(":1232"))
}
