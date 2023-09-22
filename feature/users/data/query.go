package data

import (
	"errors"
	"fmt"
	"hris-app-golang/feature/users"
	"hris-app-golang/helper"
	"mime/multipart"

	"gorm.io/gorm"
)

type UserQuery struct {
	db        *gorm.DB
	dataLogin users.UserCore
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserQuery{
		db: db,
	}
}

// GetAllManager implements users.UserDataInterface.
func (repo *UserQuery) GetAllManager() ([]users.UserCore, error) {
	var userModel []User
	tx := repo.db.Preload("Division").Where("role_id = 3").Find(&userModel)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("no row affected")
	}
	var userCore []users.UserCore
	for _, value := range userModel {
		var user = ModelToCore(value)
		userCore = append(userCore, user)
	}
	return userCore, nil
}

// Insert implements users.UserDataInterface.
func (repo *UserQuery) Insert(input users.UserCore, file multipart.File, fileName string) error {
	var userModel = UserCoreToModel(input)
	var userLead User

	hass, errHass := helper.HassPassword(userModel.Password)
	if errHass != nil {
		return errHass
	}
	userModel.Password = hass
	if userModel.RoleID == 1 {
		userModel.DivisionID = 1
	} else if userModel.RoleID == 2 {
		userModel.DivisionID = 2
	} else {
		repo.db.Where("id = ?", userModel.UserLeadID).First(&userLead)
		userModel.DivisionID = userLead.DivisionID
	}

	if fileName == "default.jpg" {
		userModel.ProfilePhoto = fileName
	} else {
		nameGen, errGen := helper.GenerateName()
		if errGen != nil {
			return errGen
		}
		userModel.ProfilePhoto = nameGen + fileName
		errUp := helper.Uploader.UploadFile(file, userModel.ProfilePhoto)

		if errUp != nil {
			return errUp
		}
	}

	tx := repo.db.Create(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}

	return nil
}

func (repo *UserQuery) SelectById(id uint) (users.UserCore, error) {
	var result User
	tx := repo.db.Preload("Role").Preload("Division").Preload("UserImport").Find(&result, id)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}

	resultCore := ModelToCore(result)
	return resultCore, nil
}

func (repo *UserQuery) Delete(id uint) error {
	tx := repo.db.Where("id = ?", id).Delete(&User{})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

// SelectAll implements users.UserDataInterface.
func (repo *UserQuery) SelectAll(role_id uint, division_id, page, item uint, search_name string) ([]users.UserCore, int64, error) {
	var userModel []User
	var count int64
	var tx *gorm.DB
	var query = repo.db.Preload("Role").Preload("Division").Preload("UserImport")

	if role_id == 3 || role_id == 4 {
		query = query.Where("role_id in (3,4) and division_id = ?", division_id)
	} else if role_id == 2 {
		query = query.Where("role_id in (3,4)")
	} else if role_id == 1 {
		query = query.Where("role_id in (1,2,3,4)")
	}

	if search_name != "" {
		query = query.Where("first_name like ? or last_name like ?", "%"+search_name+"%", "%"+search_name+"%")
	}

	queryCount := query
	tx = queryCount.Find(&userModel)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, 0, errors.New("no row affected")
	}
	count = tx.RowsAffected

	if page != 0 && item != 0 {
		limit := item
		offset := (page - 1) * item
		query = query.Limit(int(limit)).Offset(int(offset))
		fmt.Println(limit, offset)
	}
	tx = query.Find(&userModel)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, 0, errors.New("no row affected")
	}

	var userCore []users.UserCore

	for _, value := range userModel {
		var user = ModelToCore(value)
		userCore = append(userCore, user)
	}

	return userCore, count, nil
}

// Update implements users.UserDataInterface.
func (repo *UserQuery) Update(id uint, input users.UserCore, file multipart.File, fileName string) error {
	var userModel = UserCoreToModel(input)
	var userLead User

	repo.db.Where("id = ?", userModel.UserLeadID).First(&userLead)
	userModel.DivisionID = userLead.DivisionID

	hass, errHass := helper.HassPassword(userModel.Password)
	if errHass != nil {
		return errHass
	}
	userModel.Password = hass

	if fileName == "default.jpg" {
		userModel.ProfilePhoto = fileName
	} else {
		nameGen, errGen := helper.GenerateName()
		if errGen != nil {
			return errGen
		}
		userModel.ProfilePhoto = nameGen + fileName
		errUp := helper.Uploader.UploadFile(file, userModel.ProfilePhoto)

		if errUp != nil {
			return errUp
		}
	}

	tx := repo.db.Session(&gorm.Session{FullSaveAssociations: true}).Model(&User{}).Where("id = ?", id).Updates(&userModel)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("no row affected")
	}
	return nil
}

func (repo *UserQuery) Login(email string, password string) (dataLogin users.UserCore, err error) {

	var data User

	tx := repo.db.Where("email = ?", email).Preload("Role").Preload("Division").Find(&data)
	if tx.Error != nil {
		return users.UserCore{}, tx.Error
	}
	check := helper.CheckPassword(password, data.Password)
	if !check {
		return users.UserCore{}, errors.New("password incorect")
	}
	if tx.RowsAffected == 0 {
		return users.UserCore{}, errors.New("no row affected")
	}
	dataLogin = ModelToCore(data)
	repo.dataLogin = dataLogin
	return dataLogin, nil
}

// CountEmployees menghitung jumlah employee berdasarkan jumlah ID pada tabel User.
func (repo *UserQuery) GetDashboard() (users.DashboardCore, error) {
	var employeeCount, managerCount, maleUsers, femaleUsers int64
	tx1 := repo.db.Model(&User{}).Where("role_id = 4").Count(&employeeCount)
	if tx1.Error != nil {
		return users.DashboardCore{}, tx1.Error
	}
	tx2 := repo.db.Model(&User{}).Where("role_id = 3").Count(&managerCount)
	if tx2.Error != nil {
		return users.DashboardCore{}, tx2.Error
	}
	tx3 := repo.db.Model(&User{}).Preload("UserImport", "gender = Male").Count(&maleUsers)
	if tx3.Error != nil {
		return users.DashboardCore{}, tx3.Error
	}
	tx4 := repo.db.Model(&User{}).Preload("UserImport", "gender = Female").Count(&femaleUsers)
	if tx4.Error != nil {
		return users.DashboardCore{}, tx4.Error
	}

	var dashCore = users.DashboardCore{
		EmployeeCount:   uint(employeeCount),
		ManagerCount:    uint(managerCount),
		MaleUserCount:   uint(maleUsers),
		FemaleUserCount: uint(femaleUsers),
	}

	return dashCore, nil
}
