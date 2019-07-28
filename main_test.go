package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/enbis/app-strava-server/data"
	models "github.com/enbis/app-strava-server/models"

	gonfig "github.com/tkanos/gonfig"
)

var configuration models.Configuration

func init() {
	fmt.Println("init")

	err := gonfig.GetConf("./config.json", &configuration)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	fmt.Println("token ", configuration.Token)

}

func TestMain(t *testing.T) {
	fmt.Println("-------------MainTest-----------")
	u := Prova{Id: 3}
	b := new(bytes.Buffer)

	json.NewEncoder(b).Encode(u)

	res, _ := http.Post("http://localhost:3300", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}

func TestRedirectStravaAuthPage(t *testing.T) {
	fmt.Println("-------------TestRedirectStravaAuthPage-----------")
	req, err := http.NewRequest("GET", "https://www.strava.com/oauth/authorize", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("client_id", configuration.ClientId)
	q.Add("redirect_uri", "http://localhost:3300")
	q.Add("response_type", "code")
	q.Add("scope", "activity:write")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))
}

//Refresh Token with Auth readAll
func TestRefreshToken(t *testing.T) {
	fmt.Println("-------------TestRefreshToken-----------")
	req, err := http.NewRequest("POST", "https://www.strava.com/oauth/token", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("client_id", configuration.ClientId)
	q.Add("client_secret", configuration.ClientSecret)
	q.Add("grant_type", "refresh_token")
	q.Add("refresh_token", configuration.RefreshToken)
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string([]byte(body)))
	//write on json file new token

	var refreshedToken models.RefreshedToken
	json.Unmarshal([]byte(body), &refreshedToken)

	data.OverwriteJsonToken(refreshedToken.Access_Token)

}

//last 30 activities
func TestListActs(t *testing.T) {
	fmt.Println("-------------TestListActs-----------", configuration.Token)

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
	}

	fmt.Printf("%+v", response)
	//log.Println(string([]byte(body)))

}

func TestGetAct(t *testing.T) {

	fmt.Println("-------------TestGetAct-----------", configuration.Token)

	var bearer = "Bearer " + configuration.Token

	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/activities/2567553625?include_all_efforts=", nil)
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
	log.Println(string([]byte(body)))

}

func TestOverwriteFile(t *testing.T) {
	var testconfig models.Configuration

	//OpenFile
	jsonfile, err := os.Open("test.json")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer jsonfile.Close()

	byteValues, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		log.Fatalf("failed reading json file: %s", err)
	}

	json.Unmarshal(byteValues, &testconfig)

	testconfig.Token = "newToken"

	newJsonFile, err := json.MarshalIndent(testconfig, "", "")
	if err != nil {
		log.Fatalf("failed marshalling json file: %s", err)
	}

	err = ioutil.WriteFile("test.json", newJsonFile, 0644)
	if err != nil {
		log.Fatalf("failed write new json file: %s", err)
	}
}

func TestNestedJson(t *testing.T) {

	jsonfile, err := os.Open("testNested.json")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer jsonfile.Close()

	byteValues, err := ioutil.ReadAll(jsonfile)
	if err != nil {
		log.Fatalf("failed reading json file: %s", err)
	}

	var response models.SingleActivity
	err = json.Unmarshal(byteValues, &response)

	if err == nil {
		for _, lap := range response.Laps {
			fmt.Println(lap)
		}
	} else {
		fmt.Println("error unm ", err.Error())
	}
}
