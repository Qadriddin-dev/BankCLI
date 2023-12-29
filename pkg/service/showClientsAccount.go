package service

import (
	"fmt"
	"github.com/fatih/color"
)

func (s *Service) ShowClientsAccount() {
	var name string

	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)
	client, err := s.repo.GetClientsBalance(name)

	if err != nil {
		color.Red("Ошибка!, данного пользователя нет в нашей бд")
		return
	}

	fmt.Printf("Имя: %v\nБаланс: %v\n", client.Name, client.Balance)
}
