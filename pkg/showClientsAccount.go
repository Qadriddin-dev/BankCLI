package pkg

import "fmt"

func ShowClientsAccount() {
	var name string

	fmt.Println("Введите имя клиента")
	fmt.Scan(&name)

	// проверка на наличие клиента
	clientFound := false
	for _, client := range Clients {
		if client.Name == name {
			clientFound = true

			fmt.Printf("Name: %v\nBalance: %v\n", client.Name, client.Balance)
			break
		}
	}

	if !clientFound {
		fmt.Println("Ошибка!, данного пользователя нет в нашей бд")
	}
}
