package pkg

import "fmt"

func ShowClientBalance() {
	var name string
	fmt.Println("Введите имя клиента:")
	fmt.Scan(&name)
	client := clients[name]
	if client != nil {
		fmt.Println("________________")
		fmt.Println("Баланс клиента", name, ":", client["balance"], "с.")
		fmt.Println("________________")
	} else {
		fmt.Println("________________")
		fmt.Println("Клиент не найден")
		fmt.Println("________________")
	}
}
