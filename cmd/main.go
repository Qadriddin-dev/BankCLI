package main

import (
	"bankCLI/pkg/database"
	"bankCLI/pkg/repository"
	"bankCLI/pkg/service"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	db := database.NewDatabase()

	repo := repository.NewRepository(db)

	svc := service.NewService(repo)

	svc.FillDatabase()
	for {
		var choice int

		fmt.Println("1. Создание клиента, и его счет")
		fmt.Println("2. Пополнить счет клиента")
		fmt.Println("3. Посмотреть баланс клиента")
		fmt.Println("4. Снять деньги с баланса")
		fmt.Println("5. Перевод денег")
		fmt.Println("6. Получить прибыль банка")
		fmt.Println("7. Получить общий счет определоного города")
		fmt.Println("8. Выйти")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			svc.RegisterClient()
		case 2:
			svc.ToUpClientsAccount()
		case 3:
			svc.ShowClientsAccount()
		case 4:
			svc.WithDrawClientsAccount()
		case 5:
			svc.TransferMoney()
		case 6:
			svc.ShowProfit()
		case 7:
			svc.ShowAllCityAmount()
		case 8:
			color.Green("До свидания!")
			return
		}
	}
}
