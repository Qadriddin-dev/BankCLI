package main

import (
	"fmt"
	"time"
)

var clients = make(map[string]map[string]int)

func addClient() {
	var name string
	var age int

	fmt.Println("Введите имя клиента:")
	fmt.Scan(&name)

	fmt.Println("Введите возраст клиента:")
	fmt.Scan(&age)

	client := make(map[string]int)
	client["age"] = age
	client["balance"] = 0

	clients[name] = client

	fmt.Println("________________")
	fmt.Println("Клиент добавлен:", name)
	fmt.Println("________________")
}

func showAllClients() {
	fmt.Println("Список клиентов:")
	for name, client := range clients {
		fmt.Println(name, "- возраст:", client["age"], "- баланс:", client["balance"])
	}
	fmt.Println("______________")
}

func deposit(name string, amount int) {
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

func showBalance(name string) {
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

func withDraw(name string, amount int) {
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

func main() {
	for {
		var choice int
		fmt.Println("1. Добавить клиента")
		fmt.Println("2. Пополнить счет клиента")
		fmt.Println("3. Снять деньги со счета клиента")
		fmt.Println("4. Показать всех клиентов")
		fmt.Println("5. Показать баланс клиента")
		fmt.Println("6. Выйти")

		fmt.Scan(&choice)

		if choice == 1 {
			addClient()
		} else if choice == 2 {
			var name string
			var amount int
			fmt.Println("Введите имя клиента:")
			fmt.Scan(&name)
			fmt.Println("Введите сумму для пополнения:")
			fmt.Scan(&amount)
			deposit(name, amount)
		} else if choice == 3 {
			var name string
			var amount int
			fmt.Println("Введите имя клиента:")
			fmt.Scan(&name)
			fmt.Println("Введите сумму для снятия:")
			fmt.Scan(&amount)
			withDraw(name, amount)
		} else if choice == 4 {
			showAllClients()
		} else if choice == 5 {
			var name string
			fmt.Println("Введите имя клиента:")
			fmt.Scan(&name)
			showBalance(name)
		} else if choice == 6 {
			fmt.Println("Спасибо что воспользовались нашим сервисом :)")
			break
		}

		time.Sleep(3 * time.Second)
	}
}
