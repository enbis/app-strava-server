package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/enbis/app-strava-server/data"
	"github.com/enbis/app-strava-server/stravaAPI"
	"github.com/wcharczuk/go-chart"
)

type Prova struct {
	Id int
}

func main() {
	http.HandleFunc("/piechart", requestDataPieChart)
	http.HandleFunc("/barchart", requestDataBarChart)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var p Prova

		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		fmt.Println(r.Body)
		fmt.Println(p.Id)
	})

	http.ListenAndServe(":3300", nil)
}

func requestDataBarChart(res http.ResponseWriter, req *http.Request) {

	acts, err := stravaAPI.RequestActivities()

	if err != nil {
		log.Print(err)
	} else {
		run, bike, swim := data.GetTimeOfActs(acts)

		fmt.Println(run, bike, swim)

		sbc := chart.BarChart{
			Title:      "Workout time",
			TitleStyle: chart.StyleShow(),
			Background: chart.Style{
				Padding: chart.Box{
					Top: 40,
				},
			},
			Height:   512,
			BarWidth: 100,
			XAxis:    chart.StyleShow(),
			YAxis: chart.YAxis{
				Style: chart.StyleShow(),
			},
			Bars: []chart.Value{
				{Value: run, Label: "Run"},
				{Value: bike, Label: "Bike"},
				{Value: swim, Label: "Swim"},
			},
		}

		res.Header().Set("Content-Type", "image/png")
		err := sbc.Render(chart.PNG, res)
		if err != nil {
			fmt.Printf("Error rendering chart: %v\n", err)
		}
	}
}

func requestDataPieChart(res http.ResponseWriter, req *http.Request) {

	acts, err := stravaAPI.RequestActivities()

	if err != nil {
		log.Print(err)
	} else {
		run, bike, swim := data.GetNumberOfActs(acts)
		pie := chart.PieChart{
			Width:  512,
			Height: 512,
			Values: []chart.Value{
				{Value: run, Label: fmt.Sprintf("Run = %v", run)},
				{Value: bike, Label: fmt.Sprintf("Bike = %v", bike)},
				{Value: swim, Label: fmt.Sprintf("Swim = %v", swim)},
			},
		}

		res.Header().Set("Content-Type", "image/png")
		err := pie.Render(chart.PNG, res)
		if err != nil {
			fmt.Printf("Error rendering pie chart: %v\n", err)
		}
	}
}
