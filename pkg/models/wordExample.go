package models

import (
	"gorm.io/gorm"
)

type Example struct {
	Id          int    `json:"id" gorm:"id"`
	Simplified  string `json:"simplified" gorm:"simplified""`
	Hanzi       string `json:"hanzi" gorm:"hanzi"`
	Pinyin      string `json:"pinyin" gorm:"pinyin"`
	Translation string `json:"translation" gorm:"translation"`
	Audio       string `json:"audio" gorm:"audio"`
	Level       string `json:"level" gorm:"level"`
}

func (Example) TableName() string { return "character_example" }

func InsertExample(db *gorm.DB, example Example) bool {
	if err := db.Create(&example).Error; err != nil {
		return false
	}
	return true
}
func InsertExamples(db *gorm.DB, examples []Example) bool {
	if err := db.Create(&examples).Error; err != nil {
		return false
	}
	return true
}

func QueryExample(db *gorm.DB, character string, isQueryAudio bool) ([]Example, error) {
	var examples []Example
	selectQuery := "hanzi,simplified,pinyin,translation,level"
	if isQueryAudio {
		selectQuery += ",audio"
	}
	if err := db.Select(selectQuery).Where("simplified=?", character).Find(&examples).Error; err != nil {
		return []Example{}, err
	}
	return examples, nil
}
