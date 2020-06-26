package api

import (
	"sync"
)

// Store
type Store struct {
	Quotes         map[int32]Quote
	Customers      map[int32]Customer
	NextQuoteID    int32
	NextCustomerID int32
	Lock           sync.Mutex
}

// NewStore
func NewStore(seedQuotes map[int32]Quote, seedCustomers map[int32]Customer) *Store {
	var maxQuoteID int32
	var maxCustomerID int32

	for key := range seedQuotes {
		if key > maxQuoteID {
			maxQuoteID = key
		}
	}

	for key := range seedCustomers {
		if key > maxCustomerID {
			maxCustomerID = key
		}
	}

	return &Store{
		Quotes:         seedQuotes,
		Customers:      seedCustomers,
		NextQuoteID:    maxQuoteID,
		NextCustomerID: maxCustomerID,
	}
}
