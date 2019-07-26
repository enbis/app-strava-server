package stravaAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/enbis/app-strava-server/models"
	"github.com/tkanos/gonfig"
)

var configuration models.Configuration

func init() {

	fmt.Println("init")

	err := gonfig.GetConf("./config.json", &configuration)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func RequestActivities() ([]models.Activity, models.Response) {

	var bearer = "Bearer " + configuration.Token

	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/athlete/activities", nil)
	req.Header.Add("Authorization", bearer)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bodyByte := []byte(body)
	var response models.Response
	err = json.Unmarshal(bodyByte, &response)
	if err != nil {
		activities := make([]models.Activity, 0)
		err = json.Unmarshal(bodyByte, &activities)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}
		fmt.Printf("%+v", activities)
		return activities, response
	}

	return nil, response
}
