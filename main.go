package main

import (
	"github.com/gofiber/fiber/v2/middleware/cors"
	_handler "parking-pos-backend/src/handler"
	_repo "parking-pos-backend/src/repository"
	_usecase "parking-pos-backend/src/usecase"

	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ilyakaznacheev/cleanenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"os"
	"parking-pos-backend/config"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	var cfg config.AppConfig
	if err := cleanenv.ReadConfig("./config/app-config.json", &cfg); err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	dsn := "host=" + cfg.Database.Host + " user=" + cfg.Database.Username + " password=" + cfg.Database.Password + " dbname=" + cfg.Database.Name + " port=" + cfg.Database.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	fmt.Println(dsn)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Info),
	})

	pr := _repo.NewParkingRepository(db)
	pu := _usecase.NewParkingUsecase(pr)

	_handler.NewParkingHandler(app, pu)

	app.Listen(":3000")
}
