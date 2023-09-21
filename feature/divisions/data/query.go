package data

import (
	"hris-app-golang/feature/divisions"

	"gorm.io/gorm"
)

type DivisionQuery struct {
	db *gorm.DB
}

func NewDivisionQuery(db *gorm.DB) divisions.DivisionDataInterface {
	return &DivisionQuery{
		db: db,
	}
}

func (repo *DivisionQuery) Read() ([]divisions.DivisionCore, error) {
	var divisionsData []Division
	tx := repo.db.Find(&divisionsData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var divisionsCore []divisions.DivisionCore
	for _, value := range divisionsData {
		divisionsCore = append(divisionsCore, divisions.DivisionCore{
			ID:   value.ID,
			Name: value.Name,
		})
	}
	return divisionsCore, nil
}
