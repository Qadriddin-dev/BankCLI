package service

import (
	"bankCLI/pkg/models"
	"fmt"
	"github.com/fatih/color"
)

func (s *Service) RegisterClient() {
	var name, phoneNumber, cityName string
	var year int
	var client models.Clients

	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)
	client.Name = name

	fmt.Println("Введите год рождения клиента")
	fmt.Scan(&year)
	client.BirthYear = year

	fmt.Println("Введите номер телефона клиента")
	fmt.Scan(&phoneNumber)
	client.PhoneNumber = phoneNumber

	client.Balance = 0

	for {
		fmt.Println("Выберите город из списка:")
		cities := s.repo.CityList()
		for _, city := range cities {
			fmt.Println(city.Name)
		}
		fmt.Scan(&cityName)
		city, err := s.repo.FindCity(cityName)
		if err != nil {
			color.Red("Город не найден, попробуйте еще раз!")
			continue
		}
		client.City = city
		break
	}
	s.repo.AddClients(name, phoneNumber, year, client.City)
	color.Green("Клиент успешно добавлен!")
}
