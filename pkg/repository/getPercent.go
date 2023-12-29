package repository

func (repo *Repository) GetPercent() int {
	return repo.db.Percent
}
