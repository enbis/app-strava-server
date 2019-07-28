package stravaAPI

import (
	"encoding/json"
	"errors"
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

func RequestActivities() ([]models.Activity, error) {

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
		return activities, nil
	}

	return nil, errors.New(response.Message)
}

func RequestSingleAct(idAct string) (models.SingleActivity, error) {
	var bearer = "Bearer " + configuration.Token

	full_req := fmt.Sprintf("https://www.strava.com/api/v3/activities/%s?include_all_efforts=", idAct)

	req, err := http.NewRequest("GET", full_req, nil)
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
	var activity models.SingleActivity

	err = json.Unmarshal(bodyByte, &response)

	if err != nil {
		err = json.Unmarshal(bodyByte, &activity)

		return activity, nil
	}

	return activity, errors.New(response.Message)

}
