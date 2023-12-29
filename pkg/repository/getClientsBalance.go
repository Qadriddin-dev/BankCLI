package repository

import (
	"bankCLI/pkg/models"
	"errors"
)

func (repo *Repository) GetClientsBalance(clientName string) (*models.Clients, error) {
	for _, client := range repo.db.Clients {
		if client.Name == clientName {
			return client, nil
		}
	}
	return nil, errors.New("Ошибка!, данного пользователя нет в нашей бд")
}
