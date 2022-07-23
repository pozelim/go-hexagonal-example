package main

import (
	"fmt"
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"github.com/pozelim/go-hexagonal-example/internal/app/domain"
	"github.com/pozelim/go-hexagonal-example/internal/configuration"
	"github.com/pozelim/go-hexagonal-example/internal/database"
)

func main() {
	fmt.Println("I'm alive!")
	_ = configuration.New()
	estateService := app.NewEstateService(database.NewEstateRepositoryAdapter())
	e, _ := estateService.Create(domain.Estate{
		Name: "Estate test",
		Point: domain.Point{
			Lat: 15,
			Lon: 15,
		},
		Price: domain.Price{
			Value: 50000,
		},
		Area: 180,
	})

	fmt.Printf("%++v\n", e)
}
