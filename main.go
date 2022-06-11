package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"tz_getblockio/handler"
	"tz_getblockio/service"
)

func main() {
	cli := http.Client{Timeout: time.Duration(3) * time.Second}
	svc := service.NewMainNetService(cli)
	h := handler.NewHandler(svc)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/maximum-change", h.GetMaximumChange)
	e.Logger.Fatal(e.Start(":9090"))

}
