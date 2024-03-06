package models

import "github.com/jackc/pgx/v5/pgtype"

// ---------------------------
// EMPLOYEE
// ---------------------------

type CreateEmployeeParams struct {
	UserName     string        `db:"username" json:"userName"`
	PasswordHash string        `db:"password_hash" json:"passwordHash"`
	FirstName    string        `db:"first_name" json:"firstName"`
	LastName     string        `db:"last_name" json:"lastName"`
	Email        string        `db:"email" json:"email"`
	PhonePrimary pgtype.Text   `db:"phone_primary" json:"phonePrimary"`
	PhoneOther   []pgtype.Text `db:"phone_other" json:"phoneOther"`
	EmployeeType string        `db:"employee_type" json:"employeeType"`
}

type CreateEmployeeRequest struct {
	UserName      string        `db:"username" json:"userName"`
	PasswordPlain string        `db:"password_plain" json:"passwordPlain"`
	FirstName     string        `db:"first_name" json:"firstName"`
	LastName      string        `db:"last_name" json:"lastName"`
	Email         string        `db:"email" json:"email"`
	PhonePrimary  pgtype.Text   `db:"phone_primary" json:"phonePrimary"`
	PhoneOther    []pgtype.Text `db:"phone_other" json:"phoneOther"`
	EmployeeType  string        `db:"employee_type" json:"employeeType"`
}

type EmployeeLoginRequest struct {
	UserName      string `db:"username" json:"userName"`
	PasswordPlain string `db:"password_plain" json:"passwordPlain"`
}

type GetEmployeeResponse struct {
	UserName     string        `db:"username" json:"userName"`
	FirstName    string        `db:"first_name" json:"firstName"`
	LastName     string        `db:"last_name" json:"lastName"`
	Email        string        `db:"email" json:"email"`
	PhonePrimary pgtype.Text   `db:"phone_primary" json:"phonePrimary"`
	PhoneOther   []pgtype.Text `db:"phone_other" json:"phoneOther"`
	EmployeeType string        `db:"employee_type" json:"employeeType"`
}

// ---------------------------
// CUSTOMER
// ---------------------------

type CreateCustomerParams struct {
	UserName     string        `db:"username" json:"userName"`
	PasswordHash string        `db:"password_hash" json:"passwordHash"`
	FirstName    string        `db:"first_name" json:"firstName"`
	LastName     string        `db:"last_name" json:"lastName"`
	Email        string        `db:"email" json:"email"`
	PhonePrimary pgtype.Text   `db:"phone_primary" json:"phonePrimary"`
	PhoneOther   []pgtype.Text `db:"phone_other" json:"phoneOther"`
}

type CreateCustomerRequest struct {
	UserName      string        `db:"username" json:"userName"`
	PasswordPlain string        `db:"password_plain" json:"passwordPlain"`
	FirstName     string        `db:"first_name" json:"firstName"`
	LastName      string        `db:"last_name" json:"lastName"`
	Email         string        `db:"email" json:"email"`
	PhonePrimary  pgtype.Text   `db:"phone_primary" json:"phonePrimary"`
	PhoneOther    []pgtype.Text `db:"phone_other" json:"phoneOther"`
}

type GetCustomerResponse struct {
	UserName     string        `db:"username" json:"userName"`
	FirstName    string        `db:"first_name" json:"firstName"`
	LastName     string        `db:"last_name" json:"lastName"`
	Email        string        `db:"email" json:"email"`
	PhonePrimary pgtype.Text   `db:"phone_primary" json:"phonePrimary"`
	PhoneOther   []pgtype.Text `db:"phone_other" json:"phoneOther"`
}

type CustomerLoginRequest struct {
	UserName      string `db:"username" json:"userName"`
	PasswordPlain string `db:"password_plain" json:"passwordPlain"`
}

// ---------------------------
// JOBS
// ---------------------------

type JobsDisplayRequest struct {
	Status    string `json:"status"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

type JobResponse struct {
	ID         int                    `db:"job_id" json:"id"`
	Customer   GetCustomerResponse    `json:"customer"`
	LoadAddr   Address                `json:"loadAddr"`
	UnloadAddr Address                `json:"unloadAddr"`
	StartTime  pgtype.Timestamp       `db:"start_time" json:"startTime"`
	HoursLabor pgtype.Interval        `db:"hours_labor" json:"hoursLabor"`
	Finalized  bool                   `db:"finalized" json:"finalized"`
	Rooms      map[string]interface{} `db:"rooms" json:"rooms"`
	Pack       bool                   `db:"pack" json:"pack"`
	Unpack     bool                   `db:"unpack" json:"unpack"`
	Load       bool                   `db:"load" json:"load"`
	Unload     bool                   `db:"unload" json:"unload"`
	Clean      bool                   `db:"clean" json:"clean"`
	Milage     int                    `db:"milage" json:"milage"`
	Cost       string                 `db:"cost" json:"cost"`
	//AssignedEmp []Employee             `json:"assignedEmployees"`
}