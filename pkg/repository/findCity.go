package repository

import (
	"bankCLI/pkg/models"
	"errors"
	"fmt"
)

func (repo *Repository) FindCity(cityName string) (models.City, error) {
	for _, city := range repo.db.City {
		fmt.Println(city.Name)
		if city.Name == cityName {
			return *city, nil
		}
	}
	return models.City{}, errors.New("Город не найден")
}
