package pkg

import "fmt"

func ShowAllClients() {
	fmt.Println("Список клиентов:")
	for name, client := range clients {
		fmt.Println(name, "- возраст:", client["age"], "- баланс:", client["balance"], "c")
	}
	fmt.Println("______________")
}
