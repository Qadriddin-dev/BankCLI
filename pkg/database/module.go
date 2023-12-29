package database

import "bankCLI/pkg/models"

type Database struct {
	City      []*models.City
	Bank      []*models.Bank
	Clients   []*models.Clients
	Transfers []*models.Transfers
	Percent   int
}

func NewDatabase() *Database {
	return &Database{
		City:      make([]*models.City, 0),
		Bank:      make([]*models.Bank, 0),
		Clients:   make([]*models.Clients, 0),
		Transfers: make([]*models.Transfers, 0),
		Percent:   10,
	}
}
