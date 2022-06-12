package main

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
	"time"
	"tz_getblockio/internal/config"
	"tz_getblockio/internal/handler"
	"tz_getblockio/internal/service"
)

// @title           Swagger GetBlock Test Assignment
// @version         1.0
// @description     This is a service for getting the address which balance changed the most over the last 100 blocks
// @termsOfService  http://swagger.io/terms/
// @contact.name   Madyar Turgenbaev
// @contact.email  madiar.997@gmail.com
// @BasePath  /api/v1
func initConfig() (config.Config, error) {
	var cfg config.Config
	f, err := os.Open("./configs/config.yaml")
	if err != nil {
		return cfg, err
	}
	defer f.Close()
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}

func main() {
	cfg, err := initConfig()
	if err != nil {
		log.Fatal("error initializing configs: ", err)
	}

	cli := http.Client{Timeout: time.Duration(20) * time.Second}
	svc := service.NewMainNetService(cfg, cli)
	h := handler.NewHandler(svc)
	e := echo.New()
	g := e.Group("/api/v1")
	g.GET("/maximum-change", h.GetMaximumChange)
	e.Logger.Fatal(e.Start(cfg.Server.Port))

}
