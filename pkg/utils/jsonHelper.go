package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mandarinLearningBE/pkg/models"
	"os"
)

func JsonReader() map[string]models.CharacterJson {
	jsonFile, err := os.Open("./dict.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]models.CharacterJson
	json.Unmarshal([]byte(byteValue), &result)
	return result
}
