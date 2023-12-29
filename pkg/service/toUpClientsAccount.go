package service

import (
	"fmt"
	"github.com/fatih/color"
)

func (s *Service) ToUpClientsAccount() {
	var name string
	var amount int

	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)

	client, err := s.repo.FindClients(name)
	if err != nil {
		color.Red("Клиент не найден, попробуйте еще раз!")
		return
	}

	fmt.Println("Введите сумму для пополнения")
	fmt.Scan(&amount)

	client.Balance += amount
	color.Green("Баланс успешно пополнен!")
}
