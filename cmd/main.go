package main

import (
	"fmt"
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"github.com/pozelim/go-hexagonal-example/internal/configuration"
	"github.com/pozelim/go-hexagonal-example/internal/database"
)

func main() {
	fmt.Println("I'm alive!")
	_ = configuration.New()
	estateService := app.NewEstateService(database.NewEstateRepositoryAdapter())
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

	fmt.Printf("%++v\n", e)
}
