package service

func (s *Service) FillDatabase() {
	cities := []struct {
		name   string
		region string
	}{
		{"Dushanbe", "Dushanbe"},
		{"Khujand", "Sughd"},
		{"Kulob", "Khatlon"},
		{"Qurghonteppa", "Khatlon"},
		{"Istaravshan", "Sughd"},
	}

	for _, city := range cities {
		s.repo.AddCity(city.name, city.region)
	}
}
