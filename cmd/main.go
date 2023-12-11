package main

import (
	"BankCLI/pkg"
	"fmt"
	"time"
)

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
			pkg.RegisterClient()
		} else if choice == 2 {
			pkg.AddDeposit()
		} else if choice == 3 {
			pkg.WithDraw()
		} else if choice == 4 {
			pkg.ShowAllClients()
		} else if choice == 5 {
			pkg.ShowClientBalance()
		} else if choice == 6 {
			fmt.Println("Спасибо что воспользовались нашим сервисом :)")
			break
		}

		time.Sleep(3 * time.Second)
	}
}
