package service

import (
	"hris-app-golang/feature/divisions"
)

type DivisionService struct {
	divisionData divisions.DivisionDataInterface
}

func NewDivisionService(repo divisions.DivisionDataInterface) divisions.DivisionServiceInterface {
	return &DivisionService{
		divisionData: repo,
	}
}

func (service *DivisionService) GetDiv() ([]divisions.DivisionCore, error) {
	result, err := service.divisionData.Read()
	if err != nil {
		return []divisions.DivisionCore{}, err
	}
	return result, err
}
