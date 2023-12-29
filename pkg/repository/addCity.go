package repository

import "bankCLI/pkg/models"

func (repo *Repository) AddCity(cityName, region string) {
	repo.db.City = append(repo.db.City, &models.City{
		Name:   cityName,
		Region: region,
	})
}
