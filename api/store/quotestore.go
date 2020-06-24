package api

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type QuoteStore struct {
	Quotes map[int64]Quote
	NextId int64
	Lock   sync.Mutex
}

func NewQuoteStore() *QuoteStore {
	return &QuoteStore{
		Quotes: make(map[int64]Quote),
		NextId: 1000,
	}
}

func sendQuotestoreError(ctx echo.Context, code int, message string) error {
	quoteErr := Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, quoteErr)
	return err
}

// implement all handlers in the 'ServerInterface'

// Get all quotes
// (GET /quote)
func (q *QuoteStore) GetQuotes(ctx echo.Context) error {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	var result []Quote

	for _, quote := range q.Quotes {
		result = append(result, quote)
	}

	return ctx.JSON(http.StatusOK, result)
}

// Add a new quote
// (POST /quote)
func (q *QuoteStore) AddQuote(ctx echo.Context) error {
	var newQuote Quote
	err := ctx.Bind(&newQuote)
	if err != nil {
		return sendQuotestoreError(ctx, http.StatusBadRequest, "invalid format for new Quote")
	}

	q.Lock.Lock()
	defer q.Lock.Unlock()

	var quote Quote

	quote.Id = q.NextId
	q.NextId = q.NextId + 1
	quote.City = newQuote.City
	quote.Email = newQuote.Email
	quote.FirstName = newQuote.FirstName
	quote.LastName = newQuote.LastName
	quote.Phone = newQuote.Phone
	quote.PostCode = newQuote.PostCode
	quote.Street = newQuote.Street
	quote.Suburb = newQuote.Suburb
	quote.Tags = newQuote.Tags
	quote.Windows = newQuote.Windows

	q.Quotes[quote.Id] = quote

	err = ctx.JSON(http.StatusCreated, quote)
	if err != nil {
		return err
	}

	return nil
}

// Delete a quote
// (DELETE /quote/{quoteId})
func (q *QuoteStore) DeleteQuote(ctx echo.Context, quoteId int64, param DeleteQuoteParams) error {

}

// Find quote by ID
// (GET /quote/{quoteId})
func (q *QuoteStore) GetQuoteById(ctx echo.Context, quoteId int64) error {
	return ctx.JSON(http.StatusOK, result)
}

// Update an existing quote
// (POST /quote/{quoteId})
func (q *QuoteStore) UpdateQuoteById(ctx echo.Context, quoteId int64) error {

}
