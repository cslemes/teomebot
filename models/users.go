package models

import (
	"fmt"

	"gorm.io/gorm"
)

// ConDB.AutoMigrate(&models.TwitchUser{})

type TwitchUser struct {
	UUID       string `json:"uuid" gorm:"primaryKey"`
	TwitchId   string `json:"twitch" gorm:"unique"`
	TwitchNick string `json:"twitch_nick"`
}

func (tu *TwitchUser) Create(con *gorm.DB) error {

	res := con.Create(tu)
	if res.Error != nil {
		con.Rollback()
		return res.Error
	}

	con.Commit()
	return nil
}

func (tu *TwitchUser) Update(con *gorm.DB) error {
	res := conDB.Save(tu)
	if res.Error != nil {
		res.Rollback()
		return res.Error
	}

	res.Commit()
	return nil
}

func GetUserByField(fieldName, fieldValue string, con *gorm.DB) (*TwitchUser, error) {

	user := &TwitchUser{}

	res := con.First(&user, fmt.Sprintf("%s = ?", fieldName), fieldValue)
	if res.Error != nil {
		res.Rollback()
		return nil, res.Error
	}
	res.Commit()
	return user, nil
}
