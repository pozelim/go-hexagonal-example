package http

import (
	"github.com/labstack/echo/v4"
)

func Start(estate *Estate) {
	e := echo.New()
	e.POST("/estate", estate.Create)

	e.Start(":3000")
}
