package pkg

import (
	"bankCLI/pkg/models"
	"fmt"
)

var bank models.Bank

func TransferMoney() {
	var senderName, recipientName string
	var amount int

	fmt.Println("Введите имя отправителя")
	fmt.Scan(&senderName)

	var sender *models.Clients
	for i := range Clients {
		if Clients[i].Name == senderName {
			sender = &Clients[i]
			break
		}
	}

	if sender == nil {
		fmt.Println("Отсуствует счет отправителя")
		return
	}

	fmt.Println("Введите имя получателя")
	fmt.Scan(&recipientName)

	var recipient *models.Clients
	for i := range Clients {
		if Clients[i].Name == recipientName {
			recipient = &Clients[i]
			break
		}
	}

	if recipient == nil {
		fmt.Println("Отсуствует счет получателя")
		return
	}

	fmt.Println("Введите сумму перевода")
	fmt.Scan(&amount)

	if sender.Balance < amount {
		fmt.Println("Недостаточно средств")
		return
	}

	var comission int = amount / 100 * Percent

	sender.Balance -= amount
	recipient.Balance += amount - comission

	bank.Profit += comission

	Bank = append(Bank, bank)

	Transfers = append(Transfers, models.Transfers{
		From:   *sender,
		To:     *recipient,
		Amount: amount,
	})

	fmt.Println("Успешно переведено")
}
