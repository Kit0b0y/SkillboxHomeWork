package repository

import cities "github.com/Kit0b0y/SkillboxHomeWork/NewSkillbox/Interim_certification"

type CityList interface {
	Create(city cities.CityRequest) (string, error)
	Delete(id int) error
	SetPopulation(id, population int) error
	GetFromRegion(region string) ([]string, error)
	GetFromDistrict(district string) ([]string, error)
	GetFromPopulation(start, end int) ([]string, error)
	GetFromFoundation(start, end int) ([]string, error)
	findCities(idx_type int, searchText string) []string
	GetFull(id int) (*cities.City, error)
}

type Repository struct {
	CityList
}

func NewRepository(db *DataBase) *Repository {
	return &Repository{
		CityList: NewCityListDB(db),
	}
}
