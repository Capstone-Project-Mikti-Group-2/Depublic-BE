package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/builder"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/config"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/binder"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/server"
	"github.com/Capstone-Project-Mikti-Group-2/Depublic-BE/internal/http/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	//menghubungkan postgres
	cfg, err := config.NewConfig(".env")
	checkError(err)

	db, err := buildGormDB(cfg.Postgres)
	checkError(err)

	midtransClient := initMidtrans(cfg)

	publicRoutes := builder.BuildPublicRoutes(cfg, db, midtransClient)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db, midtransClient)

	echoBinder := &echo.DefaultBinder{}
	formValidator := validator.NewFormValidator()
	customBinder := binder.NewBinder(echoBinder, formValidator)

	srv := server.NewServer(
		cfg,
		customBinder,
		publicRoutes,
		privateRoutes,
	)
	runServer(srv, cfg.Port)
	waitForShutdown(srv)
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go func() {
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}()
}

func buildGormDB(cfg config.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.User, cfg.Password, cfg.Database, cfg.Port)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func initMidtrans(cfg *config.Config) snap.Client {
	snapClient := snap.Client{}

	if cfg.MidtransConfig.MidtransEnv == "development" {
		snapClient.New(cfg.MidtransConfig.MidtransServerKey, midtrans.Sandbox)
	} else {
		snapClient.New(cfg.MidtransConfig.MidtransServerKey, midtrans.Production)
	}

	return snapClient
}
