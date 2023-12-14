package pkg

import (
	"bankCLI/pkg/models"
	"fmt"
)

func WithdrawClientsAccount() {
	var name string
	var amount int

	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)

	var client *models.Clients
	for i := range Clients {
		if Clients[i].Name == name {
			client = &Clients[i]
			break
		}
	}

	if client == nil {
		fmt.Println("Ошибка! данного пользователя нет в нашей бд")
		return
	}

	fmt.Println("Введите сумму которую хотите снять")
	fmt.Scan(&amount)

	if client.Balance < amount {
		fmt.Println("Ошибка! недостаточно средств на балансе")
		return
	}

	client.Balance -= amount

	Clients = append(Clients, *client)
}
