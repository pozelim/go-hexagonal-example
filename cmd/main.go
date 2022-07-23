package main

import (
	"fmt"
	"github.com/pozelim/go-hexagonal-example/internal/app"
	"github.com/pozelim/go-hexagonal-example/internal/app/inmemory"
	"github.com/pozelim/go-hexagonal-example/internal/app/service"
	"github.com/pozelim/go-hexagonal-example/internal/configuration"
)

func main() {
	fmt.Println("I'm alive!")
	_ = configuration.New()
	estateService := service.NewEstateService(inmemory.NewEstateStorage())
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
