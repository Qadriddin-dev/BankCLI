package service

import (
	"github.com/fatih/color"
)

func (s *Service) ShowProfit() {
	bankProfit := s.repo.GetBankProfit()
	color.Green("Прибыль банка составляет: %d", bankProfit)
}
