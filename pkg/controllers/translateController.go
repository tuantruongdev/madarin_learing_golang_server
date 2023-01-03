package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TranslateBody struct {
	Translate []string `json:"translates"`
}
type TranslateRequest struct {
	Q       string `json:"q"`
	Source  string `json:"source"`
	Target  string `json:"target"`
	Format  string `json:"format"`
	Api_key string `json:"api_Key"`
}

type TranslateText struct {
	TranslatedText string `json:"translatedText"`
	TranslatedKey  int
}

type TranslateResponse struct {
	translatedText http.Response
	key            int
	err            error
}

type TranslateResponseBody struct {
	response *http.Response
	err      error
}

//
//func (i *TranslateBody) UnmarshalJSON(b []byte) error {
//	return json.Unmarshal(b, &i.Translate)
//}

func Translate(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var translate TranslateBody
		translateRespons := make(map[int]string, 1)
		if err := context.BindJSON(&translate); err != nil {
			fmt.Println("errbind", err.Error())
			context.JSON(http.StatusOK, "something went wrong")
			return
		}
		fmt.Println(translate.Translate)

		translateChannel := make(chan TranslateText, len(translate.Translate))
		for i := 0; i < len(translate.Translate); i++ {
			go doTranslate(translate.Translate[i], i, translateChannel)
		}
		for i := 0; i < len(translate.Translate); i++ {
			temp := <-translateChannel
			translateRespons[temp.TranslatedKey] = temp.TranslatedText
		}
		responseArray := make([]string, 0)
		for i := 0; i < len(translateRespons); i++ {
			responseArray = append(responseArray, translateRespons[i])
		}

		fmt.Println(translateRespons)
		context.JSON(http.StatusOK, gin.H{"status": "ok", "data": responseArray})
	}
}

func doTranslate(q string, key int, c chan TranslateText) {
	var translated TranslateText
	reqBody := TranslateRequest{
		Q:       q,
		Source:  "zh",
		Target:  "en",
		Format:  "text",
		Api_key: "",
	}
	res := request("http://localhost:5000/translate", http.MethodPost, reqBody)
	fmt.Println("doing ", key, time.Now())
	if res.err != nil {
		fmt.Println(res.err.Error())
		return
	}
	body, err := ioutil.ReadAll(res.response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(body, &translated)
	translated.TranslatedKey = key
	c <- translated
}

func request(url string, method string, reqBody TranslateRequest) TranslateResponseBody {
	body, e := json.Marshal(reqBody)
	if e != nil {
		fmt.Errorf("parse error")
	}
	responseBody := bytes.NewBuffer(body)
	var (
		response *http.Response
		err      error
	)

	if method == http.MethodPost {
		response, err = http.Post(url, "application/json", responseBody)
	}
	//var translated TranslateText
	//json.Unmarshal(response, &translated)
	return TranslateResponseBody{
		response: response,
		err:      err,
	}
}
