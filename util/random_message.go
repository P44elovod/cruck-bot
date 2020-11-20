package util

import (
	"encoding/json"
	"log"
)

func GetRandomMessage(messages []string) (requestBody []byte) {
	randIndex := getRandomInt(len(messages))

	requestBody, err := json.Marshal(map[string]string{
		"text": messages[randIndex],
	})
	if err != nil {
		log.Fatal(err)
	}

	return
}
