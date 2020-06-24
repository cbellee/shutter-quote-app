package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cbellee/shutter-quote-app/server"
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

	quoteStore := api.NewQuoteStore()
	e := echo.New()
	e.Use((echomiddleware.Logger()))
	e.Use(middleware.OapiRequestValidator(swagger))
	api.RegisterHandlers(e, quoteStore)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
