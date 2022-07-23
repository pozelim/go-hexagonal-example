package main

import (
	"fmt"
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"github.com/pozelim/go-hexagonal-example/internal/app/model"
	"github.com/pozelim/go-hexagonal-example/internal/database"
)

func main() {
	fmt.Println("I'm alive!")
	estateService := app.NewEstateService(database.NewEstateRepositoryAdapter())
	e, _ := estateService.Create(model.Estate{
		Name: "Estate test",
		Point: model.Point{
			Lat: 15,
			Lon: 15,
		},
		Price: model.Price{
			Value: 50000,
		},
		Area: 180,
	})
	fmt.Printf("%++v\n", e)
}
