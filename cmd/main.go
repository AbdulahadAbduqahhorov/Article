package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/AbdulahadAbduqahhorov/gin/Article/api"
	"github.com/AbdulahadAbduqahhorov/gin/Article/api/handlers"
	"github.com/AbdulahadAbduqahhorov/gin/Article/config"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage/postgres"
)

func main() {
	var stg storage.StorageI
	var err error
	cfg := config.Load()
	stg, err = postgres.NewPostgres(fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase))
	if err != nil {
		panic(err)
	}
	h := handlers.NewHandler(stg)
	switch cfg.Environment {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	api.SetUpApi(router, h, cfg)

	router.Run(cfg.HTTPPort)
}
