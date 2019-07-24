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

	gonfig "github.com/tkanos/gonfig"
)

var configuration Configuration

type Configuration struct {
	Token    string
	ClientId string
}

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
	log.Println(string([]byte(body)))

}

func TestGetAct(t *testing.T) {

	fmt.Println("-------------TestGetAct-----------", configuration.Token)

	var bearer = "Bearer " + configuration.Token

	req, err := http.NewRequest("GET", "https://www.strava.com/api/v3/activities/2521280000?include_all_efforts=", nil)
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
