package pkg

import "fmt"

func TopUpClientsAccount() {
	var name string
	var amount int

	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)

	// проверка на наличие клиента
	clientFound := false
	for i, client := range Clients {
		if client.Name == name {
			clientFound = true

			fmt.Println("Введите сумму которую хотите пополнить")
			fmt.Scan(&amount)

			Clients[i].Balance += amount
			break
		}
	}

	if !clientFound {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
	}
}
