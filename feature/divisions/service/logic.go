package service
<<<<<<< feature/users
=======

import (
	"hris-app-golang/feature/divisions"
)

type DivisionService struct {
	divisionData divisions.DivisionDataInterface
}

func NewDivisionService(repo divisions.DivisionDataInterface) *DivisionService {
	return &DivisionService{
		divisionData: repo,
	}
}

func (service *DivisionService) Create(division divisions.DivisionCore) (divisions.DivisionCore, error) {
	createdDivision := divisions.DivisionCore{
		Name:      division.Name, // Jika ada lebih banyak properti, sesuaikan di sini
		CreatedAt: division.CreatedAt,
		UpdatedAt: division.UpdatedAt,
	}

	err := service.divisionData.Insert(createdDivision)
	if err != nil {
		return divisions.DivisionCore{}, err
	}
	return createdDivision, nil
}

func (service *DivisionService) GetAll() ([]divisions.DivisionCore, error) {
	divisions, err := service.divisionData.SelectAll()
	if err != nil {
		return nil, err
	}
	return divisions, nil
}

func (service *DivisionService) GetByID(id uint) (divisions.DivisionCore, error) {
	division, err := service.divisionData.SelectByID(id)
	if err != nil {
		return divisions.DivisionCore{}, err
	}
	return division, nil
}

func (service *DivisionService) Update(id uint, division divisions.DivisionCore) (divisions.DivisionCore, error) {
	err := service.divisionData.Update(id, division)
	if err != nil {
		return divisions.DivisionCore{}, err
	}
	return division, nil
}

func (service *DivisionService) Delete(id uint) error {
	err := service.divisionData.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
>>>>>>> local
