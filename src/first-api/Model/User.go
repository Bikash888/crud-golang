package Model

import (
	"first/src/first-api/Config"

	_ "github.com/go-sql-driver/mysql"
)

func SaveUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func SaveAccountInfo(account *Account) (err error) {
	if err = Config.DB.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func FindAllUser(user *[]User) (err error) {
	if err = Config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func FindUserById(user *User, id string) (err error) {

	if err = Config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	Config.DB.Save(user)
	return nil
}
func DeleteUser(user *User, id string) (err error) {

	Config.DB.Where("id=?", id).Delete(user)
	return nil
}
