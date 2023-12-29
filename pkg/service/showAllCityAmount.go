package service

import (
	"fmt"
	"github.com/fatih/color"
)

func (s *Service) ShowAllCityAmount() {
	var cityName string
	fmt.Println("Введите название города")
	fmt.Scan(&cityName)

	cityAmount, err := s.repo.GetCityAmount(cityName)
	if err != nil {
		color.Red("Ошибка! данного города нет в нашей бд")
		return
	}

	color.Green("Город: %v\nСчет: %v\n", cityName, cityAmount)
}
