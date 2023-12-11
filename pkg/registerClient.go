package pkg

import "fmt"

var clients = make(map[string]map[string]int)

func RegisterClient() {
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
