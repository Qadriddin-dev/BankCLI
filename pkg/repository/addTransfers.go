package repository

import "bankCLI/pkg/models"

func (repo *Repository) AddTransfers(from, to *models.Clients, amount int) {
	repo.db.Transfers = append(repo.db.Transfers, &models.Transfers{
		From:   *from,
		To:     *to,
		Amount: amount,
	})
}
