package http

import (
	"github.com/labstack/echo/v4"
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"net/http"
)

type Estate struct {
	port app.EstatePort
}

func NewEstate(port app.EstatePort) *Estate {
	return &Estate{port: port}
}

func (e *Estate) Create(c echo.Context) error {
	body := new(EstateCreateBody)
	if err := c.Bind(body); err != nil {
		return err
	}
	est, err := e.port.Create(body.toDomainType())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, est)
}

type EstateCreateBody struct {
	Name  string  `json:"name"`
	Lat   float32 `json:"lat"`
	Lon   float32 `json:"lon"`
	Price float64 `json:"price"`
	Area  int     `json:"area"`
}

func (e *EstateCreateBody) toDomainType() app.Estate {
	return app.Estate{
		Name: e.Name,
		Area: e.Area,
		Price: app.Price{
			Value: e.Price,
		},
		Point: app.Point{
			Lat: e.Lat,
			Lon: e.Lon,
		},
	}
}
