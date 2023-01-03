package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mandarinLearningBE/pkg/models"
	"net/http"
	"strings"
)

func SaveFavoriteCharacter(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, ok := context.Keys["user"].(models.User)
		if !ok || user.Id == 0 {
			//spamming paste xd
			context.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "you're not logged in"})
			return
		}
		//	fmt.Println(context.GetRawData())
		wordLookup := make([]models.CharacterLookup, 0)
		err := context.BindJSON(&wordLookup)
		if err != nil || len(wordLookup) < 1 {
			context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "you must contain list character!"})
			fmt.Println()
			return
		}

		models.RemovePrevFavorite(db, user)
		for i := 0; i < len(wordLookup); i++ {
			word := wordLookup[i]
			if strings.Compare(word.Simplified, "") == 0 {
				continue
			}
			syncWord := models.SyncWord{UserId: user.Id, FavoriteCharacter: word.Simplified, FavoriteType: ""}
			models.InsertFavorite(db, &syncWord)
		}
		//fmt.Println(wordLookup[0].Rank)
		context.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Sync word successfully"})
	}
}
func QueryFavorite(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, ok := context.Keys["user"].(models.User)
		if !ok || user.Id == 0 {
			//spamming paste xd
			context.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "you're not logged in"})
			return
		}
		syncWord := make([]models.SyncWord, 0)
		syncWord, err := models.QueryFavoriteList(db, user)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "jdklsjdlajdl"})
		}
		characterLookup := make([]models.CharacterLookup, 0)
		fmt.Println(syncWord)
		for i := 0; i < len(syncWord); i++ {

			lookup, err := models.QueryCharacter(db, syncWord[i].FavoriteCharacter)
			if err != nil {
				continue
			}
			entries, err := models.QueryEntry(db, syncWord[i].FavoriteCharacter)
			if err != nil {
				continue
			}

			lookup.Entries = entries

			characterLookup = append(characterLookup, lookup)
		}
		context.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Get word successfully", "data": characterLookup})

	}
}
