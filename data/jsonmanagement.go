package data

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/enbis/app-strava-server/models"
)

func OverwriteJsonToken(token string) {
	var config models.Configuration

	jsonfile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer jsonfile.Close()

	byteValues, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		log.Fatalf("failed reading json file: %s", err)
	}

	json.Unmarshal(byteValues, &config)

	fmt.Printf("Old token %v, new token %v", config.Token, token)

	config.Token = token

	newJsonFile, err := json.MarshalIndent(config, "", "")
	if err != nil {
		log.Fatalf("failed marshalling json file: %s", err)
	}

	err = ioutil.WriteFile("config.json", newJsonFile, 0644)
	if err != nil {
		log.Fatalf("failed write new json file: %s", err)
	}
}
