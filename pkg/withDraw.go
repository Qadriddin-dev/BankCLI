package pkg

import "fmt"

func WithDraw() {
	var name string
	var amount int
	fmt.Println("Введите имя клиента:")
	fmt.Scan(&name)
	fmt.Println("Введите сумму для снятия:")
	fmt.Scan(&amount)
	client := clients[name]
	if client != nil {
		if client["balance"] >= amount {
			client["balance"] -= amount
			fmt.Println("________________")
			fmt.Println("С баланса клиента", name, "снято", amount, "с.")
			fmt.Println("________________")
		} else {
			fmt.Println("________________")
			fmt.Println("Недостаточно средств на балансе клиента", name)
			fmt.Println("________________")
		}
	} else {
		fmt.Println("________________")
		fmt.Println("Клиент не найден")
		fmt.Println("________________")
	}
}
