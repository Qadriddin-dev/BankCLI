package repository

import "bankCLI/pkg/models"

func (repo *Repository) AddBankProfit(comission int) {
	repo.db.Bank = append(repo.db.Bank, &models.Bank{
		Profit: comission,
	})
}
