package repository

import (
	"errors"
)

func (repo *Repository) GetCityAmount(cityName string) (int, error) {
	for _, city := range repo.db.City {
		if city.Name == cityName {
			clientAmount := 0
			for _, client := range repo.db.Clients {
				if client.City.Name == cityName {
					clientAmount += client.Balance
				}
			}
			return clientAmount, nil
		}
	}
	return 0, errors.New("Ошибка!, данного города нет в нашей бд")
}
