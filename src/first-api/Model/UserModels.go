package Model

import "github.com/jinzhu/gorm"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	AccountID uint
}
type Account struct {
	gorm.Model
	AccountType   string `json:"account"`
	AccountNumber string
	Amount        int
	Branch        string
}

func (b *User) TableName() string {
	return "user"
}
func (a *Account) TableName() string {
	return "account"
}
