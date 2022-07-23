package main

import (
	"fmt"
	"github.com/pozelim/go-hexagonal-example/internal/app/http"
	"github.com/pozelim/go-hexagonal-example/internal/app/service"
	"github.com/pozelim/go-hexagonal-example/internal/app/storage/postgres"
	"github.com/pozelim/go-hexagonal-example/internal/configuration"
)

func main() {
	fmt.Println("I'm alive!")
	cfg := configuration.New()
	estateService := service.NewEstateService(postgres.NewEstateStorage(cfg.Postgres))
	http.Start(http.NewEstate(estateService))
}
