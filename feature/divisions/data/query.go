package data

import (
    "errors"
    "time"
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

func (repo *DivisionQuery) Insert(division divisions.DivisionCore) error {
    division.CreatedAt = time.Now()
    division.UpdatedAt = time.Now()
    tx := repo.db.Create(&division)
    if tx.Error != nil {
        return tx.Error
    }
    return nil
}

func (repo *DivisionQuery) SelectAll() ([]divisions.DivisionCore, error) {
    var divisions []divisions.DivisionCore
    tx := repo.db.Find(&divisions)
    if tx.Error != nil {
        return nil, tx.Error
    }
    return divisions, nil
}

func (repo *DivisionQuery) SelectByID(id uint) (divisions.DivisionCore, error) {
    var division divisions.DivisionCore
    tx := repo.db.First(&division, id)
    if tx.Error != nil {
        return divisions.DivisionCore{}, tx.Error
    }
    if tx.RowsAffected == 0 {
        return divisions.DivisionCore{}, errors.New("division not found")
    }
    return division, nil
}

func (repo *DivisionQuery) Update(id uint, division divisions.DivisionCore) error {
    var existingDivision divisions.DivisionCore
    tx := repo.db.First(&existingDivision, id)
    if tx.Error != nil {
        return tx.Error
    }
    if tx.RowsAffected == 0 {
        return errors.New("division not found")
    }

    division.ID = id
    division.CreatedAt = existingDivision.CreatedAt
    division.UpdatedAt = time.Now()

    tx = repo.db.Save(&division)
    if tx.Error != nil {
        return tx.Error
    }
    return nil
}

func (repo *DivisionQuery) Delete(id uint) error {
    var division divisions.DivisionCore
    tx := repo.db.Where("id = ?", id).Delete(&division)
    if tx.Error != nil {
        return tx.Error
    }
    return nil
}