package pkg

import "fmt"

func Transaction() {
	var fromClient, toClient string
	var amount int

	fmt.Println("Введите имя 1-го клиента:")
	fmt.Scan(&fromClient)
	from := clients[fromClient]

	if from != nil {
		fmt.Println("Введите имя 2-го клиента:")
		fmt.Scan(&toClient)
		to := clients[toClient]

		if to != nil {
			fmt.Println("Введите сумму для перевода:")
			fmt.Scan(&amount)

			if from["balance"] >= amount {
				from["balance"] -= amount

				to["balance"] += amount

				fmt.Println("________________")
				fmt.Println("Выполнен перевод на сумму:", amount, "с.")
				fmt.Println("________________")
			} else {
				fmt.Println("________________")
				fmt.Println("Недостаточно средств у клиента:", fromClient)
				fmt.Println("________________")
			}
		} else {
			fmt.Println("________________")
			fmt.Println("Клиент получатель не найден")
			fmt.Println("________________")
		}
	}
}
