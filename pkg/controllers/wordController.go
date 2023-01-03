package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"gorm.io/gorm"
	"mandarinLearningBE/pkg/models"
	"mandarinLearningBE/pkg/utils"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func LookupCharacter(db *gorm.DB) gin.HandlerFunc {
	//using here may affect performance !!
	jsonDb := utils.JsonReader()

	return func(context *gin.Context) {
		character := context.Param("character")
		if len(character) == 0 {
			context.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "query must not be empty"})
			return
		}
		dbCharacterLookup, err := models.QueryCharacter(db, character)
		//check if no result
		if dbCharacterLookup.Simplified == "" || err != nil {
			dbCharacterLookup = models.ConvertJsonToCharacter(jsonDb, character)
			////fake
			//entryLookup := Entry{Simplified: character, Traditional: character, Pinyin: "dǎ", Definitions: []string{"to beat", "to strike", "to hit", "to break", "to type", "to mix up", "to build", "to fight", "to fetch", "to make", "to tie up", "to issue"}}
			//dbCharacterLookup = CharacterLookup{
			//	Simplified: character,
			//	Rank:       43,
			//	Hsk:        1,
			//	Entries:    []Entry{entryLookup, entryLookup},
			//}
			if len(dbCharacterLookup.Entries) < 1 {
				context.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Not found"})
				return
			}
			checkHsk(db, dbCharacterLookup)
		} else {
			//get entry in case have result
			dbCharacterLookup.Entries, _ = models.QueryEntry(db, character)
		}
		//	time.Sleep(5000)
		context.JSON(http.StatusOK, gin.H{"simplified": dbCharacterLookup.Simplified, "rank": dbCharacterLookup.Rank, "hsk": dbCharacterLookup.Hsk, "entries": dbCharacterLookup.Entries})
	}
}

func checkHsk(db *gorm.DB, dbCharacterlookup models.CharacterLookup) {
	hskLevel := 1
	for i := 0; i < 6; i++ {
		if strings.Contains(utils.Hsk[i], dbCharacterlookup.Simplified) {
			hskLevel = i + 1
			break
		}
	}
	dbCharacterlookup.Hsk = hskLevel
	saveCharacterData(db, dbCharacterlookup)
}

func saveCharacterData(db *gorm.DB, dbCharacterLookup models.CharacterLookup) {
	models.InsertCharacter(db, dbCharacterLookup)
	models.InsertEntries(db, dbCharacterLookup.Entries)
}

func LookupExample(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		character := context.Param("character")
		audioQuery := context.Query("includeAudio")
		isQueryAudio := len(audioQuery) > 0 && audioQuery == "1"
		examples := make(map[string]models.Example)
		examplesArr := make([]models.Example, 0)
		audios := make([]string, 0)
		collector := colly.NewCollector()

		colly.AllowedDomains("https://www.chinesepod.com/", "chinesepod.com")
		collector.OnHTML(".sample-sentence-card div", func(element *colly.HTMLElement) {
			hanzi := element.ChildText(".simplified-sentence")
			pinyin := element.ChildText(".pinyin-sentence")
			translation := element.ChildText(".english-sentence")
			level := element.ChildText(".btn-default")
			if len(hanzi) < 1 || len(pinyin) < 1 || len(translation) < 1 || len(level) < 1 {
				return
			}
			tempExample := models.Example{Simplified: character, Hanzi: hanzi, Pinyin: pinyin, Translation: translation, Level: level}
			examples[tempExample.Hanzi] = tempExample
		})
		collector.OnHTML("script", func(element *colly.HTMLElement) {
			pattern, err := regexp.Compile("audioUrl:\"(.*?)\",")
			if err != nil {
				return
			}
			result := pattern.FindAllString(element.Text, -1)
			if len(result) < 1 {
				return
			}
			for i := 0; i < len(result); i++ {
				result[i] = strings.Replace(result[i], "audioUrl:\"", "", -1)
				result[i] = strings.Replace(result[i], "\"", "", -1)
				result[i] = strings.Replace(result[i], ",", "", -1)
				result[i], _ = strconv.Unquote(`"` + result[i] + `"`)
			}
			audios = result
		})
		collector.OnRequest(func(request *colly.Request) {
			//fmt.Println("Visiting", request.URL.String())
		})

		examplesArr, err := models.QueryExample(db, character, isQueryAudio)
		if err != nil || len(examplesArr) < 1 {
			collector.Visit("https://www.chinesepod.com/dictionary/" + character)
			i := 0
			for _, v := range examples {
				v.Audio = audios[i]
				examplesArr = append(examplesArr, v)
				i++
			}
			models.InsertExamples(db, examplesArr)
		}
		context.JSON(http.StatusOK, examplesArr)
	}

}
func AudioExample(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		//TODO update this shit
		context.JSON(http.StatusOK, gin.H{})
	}
}
