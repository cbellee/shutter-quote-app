// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get all customers
	// (GET /customer)
	GetCustomers(ctx echo.Context) error
	// Add a new customer
	// (POST /customer)
	AddCustomer(ctx echo.Context) error
	// Delete a customer
	// (DELETE /customer/{customerId})
	DeleteCustomer(ctx echo.Context, customerId int32, params DeleteCustomerParams) error
	// Find customer by ID
	// (GET /customer/{customerId})
	GetCustomerByID(ctx echo.Context, customerId int32) error
	// Update an existing customer
	// (POST /customer/{customerId})
	UpdateCustomerByID(ctx echo.Context, customerId int32) error
	// Get all quotes
	// (GET /quote)
	GetQuotes(ctx echo.Context) error
	// Add a new quote
	// (POST /quote)
	AddQuote(ctx echo.Context) error
	// Delete a quote
	// (DELETE /quote/{quoteId})
	DeleteQuote(ctx echo.Context, quoteId int32, params DeleteQuoteParams) error
	// Find quote by ID
	// (GET /quote/{quoteId})
	GetQuoteByID(ctx echo.Context, quoteId int32) error
	// Update an existing quote
	// (POST /quote/{quoteId})
	UpdateQuoteByID(ctx echo.Context, quoteId int32) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetCustomers converts echo context to params.
func (w *ServerInterfaceWrapper) GetCustomers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCustomers(ctx)
	return err
}

// AddCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) AddCustomer(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddCustomer(ctx)
	return err
}

// DeleteCustomer converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteCustomer(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "customerId" -------------
	var customerId int32

	err = runtime.BindStyledParameter("simple", false, "customerId", ctx.Param("customerId"), &customerId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter customerId: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteCustomerParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "api_key" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("api_key")]; found {
		var ApiKey string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for api_key, got %d", n))
		}

		err = runtime.BindStyledParameter("simple", false, "api_key", valueList[0], &ApiKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter api_key: %s", err))
		}

		params.ApiKey = &ApiKey
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteCustomer(ctx, customerId, params)
	return err
}

// GetCustomerByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetCustomerByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "customerId" -------------
	var customerId int32

	err = runtime.BindStyledParameter("simple", false, "customerId", ctx.Param("customerId"), &customerId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter customerId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetCustomerByID(ctx, customerId)
	return err
}

// UpdateCustomerByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateCustomerByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "customerId" -------------
	var customerId int32

	err = runtime.BindStyledParameter("simple", false, "customerId", ctx.Param("customerId"), &customerId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter customerId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateCustomerByID(ctx, customerId)
	return err
}

// GetQuotes converts echo context to params.
func (w *ServerInterfaceWrapper) GetQuotes(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetQuotes(ctx)
	return err
}

// AddQuote converts echo context to params.
func (w *ServerInterfaceWrapper) AddQuote(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AddQuote(ctx)
	return err
}

// DeleteQuote converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteQuote(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "quoteId" -------------
	var quoteId int32

	err = runtime.BindStyledParameter("simple", false, "quoteId", ctx.Param("quoteId"), &quoteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter quoteId: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params DeleteQuoteParams

	headers := ctx.Request().Header
	// ------------- Optional header parameter "api_key" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("api_key")]; found {
		var ApiKey string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for api_key, got %d", n))
		}

		err = runtime.BindStyledParameter("simple", false, "api_key", valueList[0], &ApiKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter api_key: %s", err))
		}

		params.ApiKey = &ApiKey
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteQuote(ctx, quoteId, params)
	return err
}

// GetQuoteByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetQuoteByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "quoteId" -------------
	var quoteId int32

	err = runtime.BindStyledParameter("simple", false, "quoteId", ctx.Param("quoteId"), &quoteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter quoteId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetQuoteByID(ctx, quoteId)
	return err
}

// UpdateQuoteByID converts echo context to params.
func (w *ServerInterfaceWrapper) UpdateQuoteByID(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "quoteId" -------------
	var quoteId int32

	err = runtime.BindStyledParameter("simple", false, "quoteId", ctx.Param("quoteId"), &quoteId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter quoteId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.UpdateQuoteByID(ctx, quoteId)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/customer", wrapper.GetCustomers)
	router.POST("/customer", wrapper.AddCustomer)
	router.DELETE("/customer/:customerId", wrapper.DeleteCustomer)
	router.GET("/customer/:customerId", wrapper.GetCustomerByID)
	router.POST("/customer/:customerId", wrapper.UpdateCustomerByID)
	router.GET("/quote", wrapper.GetQuotes)
	router.POST("/quote", wrapper.AddQuote)
	router.DELETE("/quote/:quoteId", wrapper.DeleteQuote)
	router.GET("/quote/:quoteId", wrapper.GetQuoteByID)
	router.POST("/quote/:quoteId", wrapper.UpdateQuoteByID)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xZ32/bNhD+Vwhuj5rkuuke9LS06YYARdet2F6KYKDFs8ROIhmSimMY/t8HkvotOZbT",
	"uD+2PUWRjse77z7efZZ2OBGFFBy40TjeYZ1kUBB3SShVoN2lVEKCMgzcfwkzW/vXbCXgGGujGE/xPsBS",
	"aJMICpMPtVEAZvpRuSrVauLRPsAKbkumgOL4Q+2iWRD4UDob3+wDnJTaiALUOPBORsxA4S6+V7DGMf4u",
	"anGIKhCi2n4f1IGJ1UdIjL0BBWH5ZDZrprThpJiGgVF7G+5JIXPA8bPFYhHgtVAFMTjGjJvnS9xsx7iB",
	"FJRdl5MHnMpMcDiOn4+5G2HHb9DAU/u7mUpbKTGBbF31GYkUoDVJZ0TrfLb2trS3pTAwsXtV8Wva8drZ",
	"8rGYG5LOJ4shaYcoRClimXlfOJL4ylmHOMAbRaS0ORpVgs16wzgVG90LEuNg3r5+8WjrEZotRqO6DsP0",
	"MO89AmO4PZwz8DvA1/2x/R2UDS7jADJgaWb6NV3amk6wjRhQjOR9bDdC0Dbi9hzVAbeWb0TJU0A+EnQ5",
	"uUgYGNSOIAo6UUwaJjgSa2Sy2seUB0k49CN8PpWLVCzph/fix0V48aKx5WWx8qYbRk3WM11eTAE0IEkF",
	"bAe2CpM6xjqIeodxi7Au4d6A4iS/EokDpgMGjvHPjFMkSoMKoQCRlb18vyGpDSjApcpxjDNjZBxF2t8O",
	"mXDHmK+FbzbckMTX33dhnKwgzwF+KliihBZrEyaisGv6W/9miY2IlOjy3bVtfiwBrqFDvUtJkgzQMlwM",
	"Y9lsNiFxT0Oh0qhaqqM3169ev33/+odluAgzU+TuJIIq9K/r96DuXMnG+UTOJLJsYMZxpoIAuRhtE74D",
	"pX3Yz8JFuLB+hQROJMMxfh4uwmeuLCZzGEfdsZf6QdtP/ncwpeIakTxHtbHdx54sYm1s+8S/gHnVeahA",
	"S2HztP6Wi0WNP3C3A5EyZ4lbHX3UdptaQszum03cUx2s67/qEU/nfkQPXSYJaL0uc9SAYtdd+Lz7xm9F",
	"iyJai5JTb3oxz9TJnqIgautBH5XFD5+2cbsBaIWO9d+v2SWldc2wP9CgzUtBtydVax6Kx4oy08sI+zoB",
	"5BsJMhkxiANQjYxAK0CEUqD22vZSbYQC3G1ebpZ+Il+/LAI508Yl17LA2axJmZsny8MLuE9KonYxyiAF",
	"DoolyBmguhQDrl9SigjisGnynCb7Pmh7WrRr5cven68cvBLsn4Qrd79zGCRRpABjwYw/7DCzYWZAqHtY",
	"9Xwi2V9/g1Vrbf4jybI7xFfmSFkFFPgdbFdu/Xek15Cx3S2PCqr9zYDfk33pmt+RnFF0fYV0aUsMhxtT",
	"kwQXZrIxeTwRmaqVU4nuPhS+Oz08dpBmPM2h6+rg6Hm5vb4aV2+Q6JVVVrU3WwTltvrMRfi2msxJI+4U",
	"KnlpdYhHTvM1pVptkSvvCSPuD0lJe7AfQ4/hPCmdR/o52PL/PD5jv7KmL8amf1qnDhME9wn4231Sek4h",
	"Yi2YNoynB9tcNY+a9w9HBfZtreJHLa7R9+eX1tXv+LPo6kO+P1lUe+SOK+q+3aScboowrOUDQtqV50wq",
	"un2v8ugjW7s48NP2KxXP50y8kc1Vtf/dmvm2oufB7hTt3J9ZOrnm+tlEsiflcYVchfy1yeOHNU2jjQ/V",
	"ZK4grtdPj4r5Wsf5OaaDzwP1t9IsvqT89eUZat9jI8mLlMcx4STJ+4TE+I9Pzqcn05Nq3IMjxC4EdVcT",
	"q337rePIwxd2XmDfLbFt+L135JNGN81Wgw9Mw0K8A2UJptujqZHg7QvxEz4rPPQhYX+z/ycAAP//mQre",
	"3/YeAAA=",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}

