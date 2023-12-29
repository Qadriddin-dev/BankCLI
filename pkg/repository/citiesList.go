package repository

import "bankCLI/pkg/models"

func (repo *Repository) CityList() []*models.City {
	return repo.db.City
}
