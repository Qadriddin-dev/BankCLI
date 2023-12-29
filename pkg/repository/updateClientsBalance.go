package repository

func (repo *Repository) UpdateClientsBalance(clientName string, newBalance int) {
	for i, client := range repo.db.Clients {
		if client.Name == clientName {
			repo.db.Clients[i].Balance = newBalance
		}
	}
}
