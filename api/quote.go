package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// SeedQuotes
func SeedQuotes(seedFileName string) map[int32]Quote {
	seedQuotes := make(map[int32]Quote)
	jsonFile, err := os.Open(seedFileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("successfully loaded %s\n", seedFileName)
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var quotes []Quote
	json.Unmarshal(byteValue, &quotes)

	for _, quote := range quotes {
		fmt.Printf("quote Id: %d\n", *quote.Id)
		seedQuotes[*quote.Id] = quote
	}
	return seedQuotes
}

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendQuotestoreError(ctx echo.Context, code int, message string) error {
	quoteErr := Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, quoteErr)
	return err
}

// GetQuotes returns all quotes in the store
func (q *Store) GetQuotes(ctx echo.Context) error {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	var result []Quote

	for _, quote := range q.Quotes {
		result = append(result, quote)
	}

	return ctx.JSON(http.StatusOK, result)
}

// AddQuote adds a new quote to the store
func (q *Store) AddQuote(ctx echo.Context) error {
	var newQuote Quote
	err := ctx.Bind(&newQuote)
	if err != nil {
		return sendQuotestoreError(ctx, http.StatusBadRequest, "Invalid format for Quote")
	}

	q.Lock.Lock()
	defer q.Lock.Unlock()

	newQuote.Id = &q.NextQuoteID
	q.NextQuoteID = q.NextQuoteID + 1
	q.Quotes[*newQuote.Id] = newQuote
	fmt.Printf("newQuote.Id: %d\n", *newQuote.Id)

	err = ctx.JSON(http.StatusCreated, newQuote)
	if err != nil {
		return err
	}
	return nil
}

// DeleteQuote removes a quote form the store
func (q *Store) DeleteQuote(ctx echo.Context, quoteId int32, params DeleteQuoteParams) error {
	q.Lock.Lock()
	defer q.Lock.Lock()

	_, found := q.Quotes[quoteId]
	if !found {
		return sendQuotestoreError(ctx, http.StatusNotFound, fmt.Sprintf("Couldn't find Quote with ID %d\n", quoteId))
	}
	delete(q.Quotes, quoteId)
	return ctx.NoContent(http.StatusNoContent)
}

// GetQuoteById gets a quote by its 'Id' field
func (q *Store) GetQuoteByID(ctx echo.Context, quoteId int32) error {
	quote, found := q.Quotes[quoteId]
	if !found {
		return sendQuotestoreError(ctx, http.StatusNotFound, fmt.Sprintf("Couldn't find Quote with ID %d\n", quoteId))
	}
	fmt.Printf("Found quote Id: %d\n", *quote.Id)
	return ctx.JSON(http.StatusOK, quote)
}

// UpdateQuoteByID updates a quote by its 'Id' field
func (q *Store) UpdateQuoteByID(ctx echo.Context, quoteId int32) error {
	var quoteToUpdate Quote
	err := ctx.Bind(&quoteToUpdate)
	if err != nil {
		return sendQuotestoreError(ctx, http.StatusBadRequest, "Invalid format for Quote\n")
	}

	q.Lock.Lock()
	defer q.Lock.Lock()

	_, found := q.Quotes[quoteId]
	if !found {
		return sendQuotestoreError(ctx, http.StatusNotFound, fmt.Sprintf("Couldn't find Quote with ID %d\n", quoteId))
	}
	q.Quotes[quoteId] = quoteToUpdate
	return ctx.JSON(http.StatusOK, quoteToUpdate)
}
