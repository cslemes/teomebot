package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func GetUserChurnProb(userId string) (float64, error) {

	url := "http://0.0.0.0:8082/churn_score/%s"
	url = fmt.Sprintf(url, userId)

	resp, err := http.Get(url)
	if err != nil {
		return 0., err
	}

	if resp.StatusCode != http.StatusOK {
		return 0., fmt.Errorf("requisição mau sucedida: %s", err)
	}

	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0., err
	}

	respBody := map[string]string{}
	if err := json.Unmarshal(bodyBytes, &respBody); err != nil {
		return 0., err
	}

	prob, err := strconv.ParseFloat(respBody["prob"], 64)
	if err != nil {
		log.Println(respBody)
		return 0., err
	}

	return prob, nil

}
