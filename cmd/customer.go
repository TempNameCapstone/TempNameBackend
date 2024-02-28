package main

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	DB "github.com/DeltaCapstone/ChoiceMoversBackend/database"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////
//Customer

// accountType must match account types ENUM in db
func getCustomer(c echo.Context) error {
	id := c.QueryParam("id")
	users, err := DB.PgInstance.GetCustomerById(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error retrieving data: %v", err))
	}
	if users == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("No user found with id: %v", id))
	}

	return c.JSON(http.StatusOK, users)
}

// POST handler to create a new user
func createCustomer(c echo.Context) error {
	var newCustomer DB.Customer
	// attempt at binding incoming json to a newUser
	if err := c.Bind(&newCustomer); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid user data"})
	}
	//validate password

	//replace plaintext password with hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(newCustomer.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Hash error: %v", err))
	}
	newCustomer.PasswordHash = string(bytes)

	// validation stuff probably needed

	userID, err := DB.PgInstance.CreateCustomer(c.Request().Context(), newCustomer)
	if err != nil || userID == 0 {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				fallthrough
			case pgerrcode.NotNullViolation:
				return c.JSON(http.StatusConflict, fmt.Sprintf("username or email already in use: %v", err))
			}
		}
		return c.JSON(http.StatusInternalServerError, fmt.Sprintf("Failed to create user: %v", err))
	}

	return c.JSON(http.StatusCreated, echo.Map{"ID": userID})
}

func updateCustomer(c echo.Context) error {
	var updatedCustomer DB.Customer
	// binding request
	if err := c.Bind(&updatedCustomer); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// update operation
	err := DB.PgInstance.UpdateCustomer(c.Request().Context(), updatedCustomer)
	if err != nil {
		// return internal server error if update fails
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update customer")
	}

	return c.JSON(http.StatusOK, "Customer updated")

}
