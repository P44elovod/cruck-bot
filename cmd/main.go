package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/P44elovod/crack-bot/util"
)

var (
	configPath string
	messages   []string = []string{
		"кря",
		"кря, бля",
		"ну йбн, кря",
		"кря, жопа в огне!!",
		"ебала жаба гадюку!",
		"та за шо?",
		"заливай!",
		"На проооод!",
		"на бля Джооон",
		"Фронтендеры...",
		"ХАйлелвел говорит: кря",
		"та сколько можно?",
		"никаких партеечек",
		"тормози лаптем",
		"туши будку",
		"не спеши, начальник",
		"три чи чотири роки",
		"А ты по русски можешь?",
		"+",
		"та ну ебать!",
	}
)

func init() {
	flag.StringVar(&configPath, "config-path", "./configs", "path to config file")

}

func main() {

	config, err := util.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	for {
		requestBody := util.GetRandomMessage(messages)
		sendCrackController(config.WebhookURL, requestBody)
		time.Sleep(time.Duration(config.Interval) * time.Minute)
	}

}

func sendCrackController(hoockURL string, message []byte) {
	resp, err := http.Post(hoockURL, "application/json", bytes.NewBuffer(message))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
}
