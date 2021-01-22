package repo

import (
	"first/src/first-api/Config"
	models "first/src/first-api/Model"
)

func GetAllUsers(user *models.User, pagination *models.Pagination) (*[]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := Config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.User{}).Where(user).Find(&users)
	if result.Error != nil {
		msg := result.Error
		return nil, totalRows, msg
	}
	errCount := Config.DB.Model(&models.User{}).Count(&totalRows).Error
	if errCount != nil {
		return nil, totalRows, errCount
	}
	return &users, totalRows, nil
}
