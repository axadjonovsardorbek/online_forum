package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"forum-service/models"
)

func GetUser(id string) (*models.User, error) {
	url := "http://auth:8088/user/" + id

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Raw response body:", string(body))

	var user models.User

	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}
