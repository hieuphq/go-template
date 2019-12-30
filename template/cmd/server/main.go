package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"{{Package}}/pkg/config"
	"{{Package}}/pkg/handler"
	"{{Package}}/pkg/logger"
	"{{Package}}/pkg/service/repo"
	"{{Package}}/pkg/service/repo/pg"
)

func main() {
	cls := config.DefaultConfigLoaders()
	cfg := config.LoadConfig(cls)
	l := initLog(cfg)

	s, close := pg.NewPostgresStore(&cfg)
	defer close()

	router := setupRouter(cfg, l, s)
	router.Run(fmt.Sprintf(":%s", cfg.Port))
}

func initLog(cfg config.Config) logger.Log {
	return logger.NewJSONLogger(
		logger.WithServiceName(cfg.ServiceName),
		logger.WithHostName(cfg.BaseURL),
	)
}

func setupRouter(cfg config.Config, l logger.Log, s repo.Store) *gin.Engine {
	r := gin.Default()

	h, err := handler.New(cfg, l, s)
	if err != nil {
		log.Fatal(err)
	}

	r.Use(cors.New(
		cors.Config{
			AllowOrigins: cfg.GetCORS(),
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
			AllowHeaders: []string{"Origin", "Host",
				"Content-Type", "Content-Length",
				"Accept-Encoding", "Accept-Language", "Accept",
				"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token"},
			ExposeHeaders:    []string{"MeAllowMethodsntent-Length"},
			AllowCredentials: true,
		},
	))

	// handlers
	r.GET("/ping", h.Ping)

	return r
}
