package main

import (
	"bufio"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type ApiResponse struct {
	Bitcoin struct {
		UAH int `json:"uah"`
	} `json:"bitcoin"`
}

func getRate(c *gin.Context) {
	var apiResponse ApiResponse

	err := GetJson(rateApiUrl, &apiResponse)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.IndentedJSON(http.StatusOK, apiResponse.Bitcoin.UAH)
}

func subscribe(c *gin.Context) {
	email := c.PostForm("email")

	if email == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	inStorage, err := emailInStorage(email)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	} 

	if inStorage {
		c.Status(http.StatusConflict)
	} else {
		if err :=addEmailToStorage(email); err != nil {
			c.Status(http.StatusInternalServerError)
			return
		} else {
			c.Status(http.StatusOK)
		}
	}
}

func sendEmailsHandler(c *gin.Context) {
	storageMutex.Lock()
	defer storageMutex.Unlock()

	file, err := os.Open(storage)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var emailArray []string
	for scanner.Scan() {
		emailArray = append(emailArray, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	var apiResponse ApiResponse

	err = GetJson(rateApiUrl, &apiResponse)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	for _, email := range emailArray {
		if govalidator.IsEmail(email){
			wg.Add(1)
			go sendEmail(email, apiResponse.Bitcoin.UAH)
		}
	}
	wg.Wait()

	c.Status(http.StatusOK)
}
