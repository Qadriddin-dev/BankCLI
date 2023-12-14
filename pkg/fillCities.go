package pkg

import "bankCLI/pkg/models"

func FillCities() {
	Cities = append(Cities, models.City{
		Name:   "Dushanbe",
		Region: "Dushanbe",
	})
	Cities = append(Cities, models.City{
		Name:   "Khujand",
		Region: "Sughd",
	})
	Cities = append(Cities, models.City{
		Name:   "Kulob",
		Region: "Khatlon",
	})
	Cities = append(Cities, models.City{
		Name:   "Qurghonteppa",
		Region: "Khatlon",
	})
	Cities = append(Cities, models.City{
		Name:   "Istaravshan",
		Region: "Sughd",
	})
}
