package models

import (
	"gorm.io/gorm"
	"time"
)

type CharacterLookup struct {
	Id         int        `json:"id" gorm:"id"`
	Simplified string     `json:"simplified" gorm:"simplified"`
	Rank       int        `json:"rank" gorm:"rank"`
	Hsk        int        `json:"hsk" gorm:"hsk"`
	Entries    []Entry    `json:"entries" gorm:"-"`
	CreatedAt  *time.Time `json:"create_at" gorm:"created_at"`
	UpdatedAt  *time.Time `json:"update_at" gorm:"updated_at"`
}

func (CharacterLookup) TableName() string { return "character_lookup" }

type Entry struct {
	Id          int        `json:"id" gorm:"id"`
	Simplified  string     `json:"simplified" gorm:"simplified"`
	Traditional string     `json:"traditional" gorm:"traditional"`
	Pinyin      string     `json:"pinyin" gorm:"pinyin"`
	Definitions []string   `json:"definitions" gorm:"serializer:json"`
	CreatedAt   *time.Time `json:"create_at" gorm:"created_at"`
	UpdatedAt   *time.Time `json:"update_at" gorm:"updated_at"`
}

func (Entry) TableName() string { return "entry" }

/*db*/
func InsertCharacter(db *gorm.DB, char CharacterLookup) bool {
	if err := db.Create(&char).Error; err != nil {
		return false
	}
	return true
}

func QueryCharacter(db *gorm.DB, character string) (CharacterLookup, error) {
	var characterLookup CharacterLookup
	if err := db.Where("simplified=?", character).First(&characterLookup).Error; err != nil {
		return CharacterLookup{}, err
	}
	return characterLookup, nil
}

func InsertEntries(db *gorm.DB, entries []Entry) bool {
	if err := db.Create(&entries).Error; err != nil {
		return false
	}
	return true
}

func InsertEntry(db *gorm.DB, entry Entry) bool {
	if err := db.Create(&entry).Error; err != nil {
		return false
	}
	return true
}

func QueryEntry(db *gorm.DB, character string) ([]Entry, error) {
	var entries []Entry
	if err := db.Where("simplified=?", character).Find(&entries).Error; err != nil {
		//fmt.Errorf(err.Error())
		return []Entry{}, err
	}
	return entries, nil
}

/*end db*/

func ConvertJsonToCharacter(jsonDb map[string]CharacterJson, character string) (characterLookup CharacterLookup) {
	characterJson := jsonDb[character]
	characterLookup.Simplified = characterJson.Simple
	characterLookup.Hsk = characterJson.Hsk
	characterLookup.Rank = characterJson.Rank

	for i := 0; i < len(characterJson.Entries); i++ {
		var tempEntry Entry
		var tempJsonEntry = characterJson.Entries[i]
		tempEntry.Traditional = tempJsonEntry.Traditional
		tempEntry.Simplified = character
		tempEntry.Pinyin = tempJsonEntry.Pinyin
		tempEntry.Definitions = tempJsonEntry.Defs
		characterLookup.Entries = append(characterLookup.Entries, tempEntry)
	}
	return
}
