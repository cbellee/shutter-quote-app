package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cbellee/shutter-quote-app/api"

	"github.com/cbellee/shutter-quote-app/api/config"
	"github.com/cbellee/shutter-quote-app/api/db"
	customerRepository "github.com/cbellee/shutter-quote-app/api/repository"
	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	var port = flag.Int("port", 8080, "Port for HTTP server test")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil
	conf, err := config.LoadConfig()
	dbClient, err := db.Connect(conf)
	c := customerRepository.NewCustomerRepository(dbClient)
	//store := api.NewStore(api.SeedQuotes("quotes.json"), api.SeedCustomers("customers.json"))

	e := echo.New()
	e.Use((echomiddleware.Logger()))
	e.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(e, store)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
