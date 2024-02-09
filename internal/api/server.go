package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func RandomStatus(q int) []int {
	time.Sleep(5 * time.Second)
	rand.Seed(time.Now().UnixNano())
	var res []int
	for i := 0; i < q; i++ {
		res = append(res, rand.Intn(200))
	}
	return res
}

type PK struct {
	PK       int `json:"id"`
	Quantity int `json:"quantity"`
}
type Result struct {
	Losses []int `json:"Losses"`
	Key    int
	Id     int
}

func PerformPUTRequest(url string, data Result) (*http.Response, error) {
	// Сериализация структуры в JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Создание PUT-запроса
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	// Выполнение запроса
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return resp, nil
}

func SendStatus(pk PK, url string) {
	// Выполнение расчётов с randomStatus
	result := RandomStatus(pk.Quantity)

	// Отправка PUT-запроса к основному серверу
	data := Result{Losses: result, Key: 123456, Id: pk.PK}
	_, err := PerformPUTRequest(url, data)
	if err != nil {
		fmt.Println("Error sending status:", err)
		return
	}

	fmt.Println("Status sent successfully for pk:", pk)
}
