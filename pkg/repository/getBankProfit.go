package repository

func (repo *Repository) GetBankProfit() int {
	totalProfit := 0
	for _, bank := range repo.db.Bank {
		totalProfit += bank.Profit
	}
	return totalProfit
}
