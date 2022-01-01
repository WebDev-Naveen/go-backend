package config

import (
	"github.com/WebDev-Naveen/go-backend/model"

	"github.com/jinzhu/gorm"
)

//OrderRepository --> Repository for Order Model
type CoinRepository interface {
	UserCoin(int) error
}

type coinRepository struct {
	connection *gorm.DB
}

//NewOrderRepository --> returns new coin repository
func NewCoinRepository() CoinRepository {
	return &coinRepository{
		connection: DB(),
	}
}

func (db *coinRepository) UserCoin(userID int) error {
	return db.connection.Create(&model.Coin{
		UserID: uint(userID),
	}).Error

}
