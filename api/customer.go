package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// This function wraps sending of an error in the Error format, and
// handling the failure to marshal that.
func sendCustomerstoreError(ctx echo.Context, code int, message string) error {
	customerErr := Error{
		Code:    int32(code),
		Message: message,
	}
	err := ctx.JSON(code, customerErr)
	return err
}

// AddCustomer adds a new customer to the store
func (c *Store) AddCustomer(ctx echo.Context) error {
	var newCustomer Customer
	err := ctx.Bind(&newCustomer)
	if err != nil {
		return sendCustomerstoreError(ctx, http.StatusBadRequest, "Invalid format for Customer")
	}

	//c.Lock.Lock()
	//defer c.Lock.Unlock()

	result, err := c.customerRepository.Insert(newCustomer)
	if err != nil {
		return err
	}

	//newCustomer.Id = &c.NextCustomerID
	//c.NextCustomerID = c.NextCustomerID + 1
	//c.Customers[*newCustomer.Id] = newCustomer
	//fmt.Printf("newCustomer.Id: %d\n", *newCustomer.Id)

	err = ctx.JSON(http.StatusCreated, newCustomer)
	if err != nil {
		return err
	}
	return nil
}

// GetCustomers returns all customers in the store
func (c *Store) GetCustomers(ctx echo.Context) error {
	c.Lock.Lock()
	defer c.Lock.Unlock()

	var result []Customer

	for _, customer := range c.Customers {
		result = append(result, customer)
	}

	return ctx.JSON(http.StatusOK, result)
}

// GetCustomerById gets a customer by its 'Id' field
func (c *Store) GetCustomerByID(ctx echo.Context, customerId int32) error {
	customer, found := c.Customers[customerId]
	if !found {
		return sendCustomerstoreError(ctx, http.StatusNotFound, fmt.Sprintf("Couldn't find Customer with ID %d", customerId))
	}
	//fmt.Printf("Found customer Id: %d\n", *customer.Id)
	return ctx.JSON(http.StatusOK, customer)
}

// UpdateCustomerById updates a customer by its 'Id' field
func (c *Store) UpdateCustomerByID(ctx echo.Context, customerId int32) error {
	var customerToUpdate Customer
	err := ctx.Bind(&customerToUpdate)
	if err != nil {
		return sendCustomerstoreError(ctx, http.StatusBadRequest, "Invalid format for Customer")
	}

	c.Lock.Lock()
	defer c.Lock.Lock()

	_, found := c.Customers[customerId]
	if !found {
		return sendCustomerstoreError(ctx, http.StatusNotFound, fmt.Sprintf("Couldn't find Customer with ID %d", customerId))
	}
	c.Customers[customerId] = customerToUpdate
	return ctx.JSON(http.StatusOK, customerToUpdate)
}

// DeleteCustomer removes a custoemr from the store
func (c *Store) DeleteCustomer(ctx echo.Context, customerId int32, params DeleteCustomerParams) error {
	c.Lock.Lock()
	defer c.Lock.Lock()

	_, found := c.Customers[customerId]
	if !found {
		return sendCustomerstoreError(ctx, http.StatusNotFound, fmt.Sprintf("Couldn't find Customer with ID %d", customerId))
	}
	delete(c.Customers, customerId)
	return ctx.NoContent(http.StatusNoContent)
}
