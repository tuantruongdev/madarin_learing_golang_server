package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mandarinLearningBE/pkg/models"
	"os"
)

var data = make(map[string]models.CharacterJson, 0)

func JsonReader() map[string]models.CharacterJson {
	if len(data) > 1 {
		return data
	}
	jsonFile, err := os.Open("./dict.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]models.CharacterJson
	json.Unmarshal([]byte(byteValue), &result)
	data = result
	return result
}
