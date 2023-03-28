package handler

import (
	"bytes"
	"cdec-api/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	grantType    = "client_credentials"
	clientID     = "EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI"
	clientSecret = "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG"
	authUrl      = "https://api.edu.cdek.ru/v2/oauth/token?parameters"
	calcUrl      = "https://api.edu.cdek.ru/v2/calculator/tarifflist"
)

func Calculate(c *gin.Context) {
	//создаем экземпляр структуры запроса
	var request models.CdecRequest
	//привязываем json к переменной request
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//get dto from request manually
	dto := models.CdecRequestDto{
		FromLocation: models.Location{Address: request.FromLocation},
		ToLocation:   models.Location{Address: request.ToLocation},
		Packages: []models.Package{
			{Weight: request.Weight,
				Lenght: request.Lenght,
				Width:  request.Width,
				Height: request.Height},
		},
	}

	//присваиваем token значение *models.jwt token
	token, err := GetAccessToken()
	if err != nil {
		fmt.Println("error getting token")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	//вызываем функцию getTraffics передавая ей значение токена и указатель на sdecRequest
	response, err := getTraffics(token, &dto)

	if err != nil {
		fmt.Println("error getting traffic data")
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	//возврашаем статус = 200  и значение response
	c.IndentedJSON(http.StatusOK, response)
}

func GetAccessToken() (*models.JwtToken, error) {
	var token models.JwtToken
	//initilization http client
	client := &http.Client{}
	data := url.Values{}
	//заполняем значения ключа значениями
	data.Set("grant_type", grantType)
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	//формируем http запрос
	req, err := http.NewRequest("POST", "https://api.edu.cdek.ru/v2/oauth/token", bytes.NewBufferString(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		return nil, err
	}

	return &token, err
}

func getTraffics(jwt *models.JwtToken, request *models.CdecRequestDto) (*models.CdecResponse, error) {
	//инициализируем экземпляр структуры
	var response models.CdecResponse
	//создаем http клиент
	client := &http.Client{}
	//преобразуем request в json
	jsonRaw, _ := json.Marshal(*request)
	//создаем http запрос и устаналиваем заголовки
	req, err := http.NewRequest("POST", calcUrl, bytes.NewBufferString(string(jsonRaw)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", jwt.AccessToken))

	if err != nil {
		return nil, err

	}
	//отправляем запрос и ждем ответ
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	//закрываем буфер
	defer resp.Body.Close()
	//заполняем данные из буфера
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response, nil
}
