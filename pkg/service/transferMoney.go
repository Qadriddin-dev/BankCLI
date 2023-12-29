package service

import (
	"fmt"
	"github.com/fatih/color"
)

func (s *Service) TransferMoney() {
	var senderName, recipientName string
	var amount int

	fmt.Println("Введите имя отправителя")
	fmt.Scan(&senderName)

	sender, err := s.repo.FindClients(senderName)

	if err != nil {
		color.Red("Отсуствует счет отправителя")
		return
	}

	fmt.Println("Введите имя получателя")
	fmt.Scan(&recipientName)

	recipient, err := s.repo.FindClients(recipientName)

	if err != nil {
		color.Red("Отсуствует счет получателя")
		return
	}

	fmt.Println("Введите сумму перевода")
	fmt.Scan(&amount)

	if sender.Balance < amount {
		color.Red("Недостаточно средств")
		return
	}

	var comission int = amount / 100 * s.repo.GetPercent()

	sender.Balance -= amount
	recipient.Balance += amount - comission

	s.repo.AddBankProfit(comission)

	s.repo.AddTransfers(sender, recipient, amount)
	color.Green("Деньги уже на счету клиента")
}
