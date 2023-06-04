package domain

import "banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepositiry interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
