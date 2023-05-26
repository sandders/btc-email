package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"net/http"
	"os"
	"strings"
	"sync"
	
)

var client *http.Client
var storageMutex = sync.Mutex{}
var wg sync.WaitGroup
var rateApiUrl = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah"
var storage = "storage.txt"

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func addEmailToStorage(email string) error {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	file, err := os.OpenFile(storage, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, email); err != nil {
		return fmt.Errorf("failed to add email to file: %v", err)
	}

	return nil
}

func emailInStorage(email string) (bool, error) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	
	if !fileExists(storage) {
		return false, nil
	}
	file, err := os.Open(storage)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == email {
			return true, nil
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading file:", err)
		return false, err
	}

	return false, nil
}


func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}

func sendEmail(address string, rate int) error {
	defer wg.Done()
	// Create a new SMTP dialer
	dialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

	// Create a new email message
	message := gomail.NewMessage()
	message.SetHeader("From", sender)
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", "BTC to UAH")
	message.SetBody("text/plain", fmt.Sprintf("1 BTS is %d UAH", rate))

	// Send the email
	return dialer.DialAndSend(message)
}
