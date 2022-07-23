package main

import (
	"fmt"
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"github.com/pozelim/go-hexagonal-example/internal/app/service"
	"github.com/pozelim/go-hexagonal-example/internal/app/storage/postgres"
	"github.com/pozelim/go-hexagonal-example/internal/configuration"
)

func main() {
	fmt.Println("I'm alive!")
	cfg := configuration.New()
	estateService := service.NewEstateService(postgres.NewEstateStorage(cfg.Postgres))
	e, _ := estateService.Create(app.Estate{
		Name: "Estate test",
		Point: app.Point{
			Lat: 15,
			Lon: 15,
		},
		Price: app.Price{
			Value: 50000,
		},
		Area: 180,
	})

	estateService.Remove(e.ID)
	fmt.Printf("%++v\n", e)
}
