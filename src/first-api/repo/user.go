package repo

import (
	"first/src/first-api/Config"
	models "first/src/first-api/Model"
)

func GetAllUsers(user *models.User, pagination *models.Pagination) (*[]models.User, error) {
	var users []models.User
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := Config.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuider.Model(&models.User{}).Where(user).Find(&users)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &users, nil
}
