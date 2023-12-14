package pkg

import (
	"bankCLI/pkg/models"
	"fmt"
)

func RegisterClient() {
	var name, cityName string
	var year, phoneNumber int
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
		for _, city := range Cities {
			fmt.Println(city.Name)
		}

		fmt.Scan(&cityName)

		cityFound := false
		for _, city := range Cities {
			if city.Name == cityName {
				client.City = city
				cityFound = true
				break
			}
		}

		if cityFound {
			break
		} else {
			fmt.Println("Пожалуйста выберите город из списка")
		}
	}
	Clients = append(Clients, client)
}
