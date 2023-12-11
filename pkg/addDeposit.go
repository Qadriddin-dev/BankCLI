package pkg

import "fmt"

func AddDeposit() {
	var name string
	var amount int
	fmt.Println("Введите имя клиента:")
	fmt.Scan(&name)
	fmt.Println("Введите сумму для пополнения:")
	fmt.Scan(&amount)
	
	client := clients[name]
	if client != nil {
		client["balance"] += amount
		fmt.Println("________________")
		fmt.Println("Баланс клиента", name, "пополнен на", amount, "с.")
		fmt.Println("________________")
	} else {
		fmt.Println("________________")
		fmt.Println("Клиент не найден")
		fmt.Println("________________")
	}
}
