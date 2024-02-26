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

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////
//Customer

// accountType must match account types ENUM in db
func getCustomers(c echo.Context) error {
	id := c.QueryParam("id")
	users, err := DB.PgInstance.GetCustomers(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error retrieving data: %v", err))
	}
	if users == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("No user found with id: %v", id))
	}

	return c.JSON(http.StatusOK, users)
}

// POST handler to create a new user
func CreateCustomer(c echo.Context) error {
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

func UpdateCustomer(c echo.Context) error {
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

///////////////////////////////////////////////////////////////////////////////////////////////////////////
//Employee

func getEmployees(c echo.Context) error {
	id := c.QueryParam("id")
	users, err := DB.PgInstance.GetEmployees(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error retrieving data: %v", err))
	}
	if users == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("No user found with id: %v", id))
	}
	return c.JSON(http.StatusOK, users)
}

func CreateEmployee(c echo.Context) error {
	var newEmployee DB.Employee

	// attempt at binding incoming json to a newUser
	if err := c.Bind(&newEmployee); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"Bind error": "Invalid user data"})
	}
	//validate password

	//replace plaintext password with hash
	bytes, err := bcrypt.GenerateFromPassword([]byte(newEmployee.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Hash error: %v", err))
	}
	newEmployee.PasswordHash = string(bytes)

	// validation stuff probably needed

	userID, err := DB.PgInstance.CreateEmployee(c.Request().Context(), newEmployee)
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

func UpdateEmployee(c echo.Context) error {
	var updatedEmployee DB.Employee

	// binding json to employee
	if err := c.Bind(&updatedEmployee); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
	}

	// update operation
	err := DB.PgInstance.UpdateEmployee(c.Request().Context(), updatedEmployee)
	if err != nil {
		// return internal server error if update fails
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update employee")
	}

	return c.JSON(http.StatusOK, "Employee updated")
}
