package service

import (
	"fmt"
	"github.com/fatih/color"
)

func (s *Service) WithDrawClientsAccount() {
	var name string
	var amount int

	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)

	client, err := s.repo.FindClients(name)
	if err != nil {
		color.Red("Ошибка!, данного пользователя нет в нашей бд")
		return
	}

	fmt.Println("Введите сумму которую хотите снять")
	fmt.Scan(&amount)

	if client.Balance < amount {
		color.Red("Ошибка! недостаточно средств на балансе")
		return
	}

	client.Balance -= amount
	newBalance := client.Balance

	color.Green("Деньги успешно сняти!")

	s.repo.UpdateClientsBalance(name, newBalance)
}
