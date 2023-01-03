package models

import (
	"fmt"
	"gorm.io/gorm"
)

type SyncWord struct {
	UserId            int    `json:"user_id" gorm:"user_id"`
	FavoriteCharacter string `json:"character_favorites" gorm:"character_favorites"`
	FavoriteType      string `json:"favorite_type" gorm:"favorite_type"`
}

func (*SyncWord) TableName() string {
	return "character_favorites"
}

func InsertFavorite(db *gorm.DB, syncWord *SyncWord) bool {
	if err := db.Create(&syncWord).Error; err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}

func RemovePrevFavorite(db *gorm.DB, user User) bool {
	if err := db.Where("user_id like ?", user.Id).Delete(&SyncWord{}).Error; err != nil {
		fmt.Println(err.Error())
		return true
	}
	return false
}

func QueryFavoriteList(db *gorm.DB, user User) ([]SyncWord, error) {
	var syncList = make([]SyncWord, 0)
	//if err := db.Where("user_id=?", user.Id).Take(&syncList).Error; err != nil {
	//	return []SyncWord{}, err
	//}
	db.Where("user_id=?", user.Id).Find(&syncList)
	return syncList, nil
}
