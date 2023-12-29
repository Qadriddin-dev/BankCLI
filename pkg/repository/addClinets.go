package repository

import (
	"bankCLI/pkg/models"
)

func (repo *Repository) AddClients(clientName, clientPhone string, clientBirthYear int, clientCity models.City) {
	repo.db.Clients = append(repo.db.Clients, &models.Clients{
		Name:        clientName,
		BirthYear:   clientBirthYear,
		PhoneNumber: clientPhone,
		City:        clientCity,
	})
}
