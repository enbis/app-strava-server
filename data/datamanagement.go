package data

import (
	"fmt"

	models "github.com/enbis/app-strava-server/models"
)

func GetNumberOfActs(acts []models.Activity) (run, bike, swim float64) {

	run, bike, swim = 0, 0, 0
	for _, act := range acts {
		fmt.Println(act.Type)
		switch act.Type {
		case "Run":
			run++
		case "Ride":
			bike++
		case "Swim":
			swim++
		}
	}

	fmt.Printf("Run %v Bike %v Swim %v", run, bike, swim)
	return

}

func GetTimeOfActs(acts []models.Activity) (run, bike, swim float64) {

	run, bike, swim = 0, 0, 0
	for _, act := range acts {
		switch act.Type {
		case "Run":
			run += float64(act.Moving_Time)
		case "Ride":
			bike += float64(act.Moving_Time)
		case "Swim":
			swim += float64(act.Moving_Time)
		}
	}

	total := run + bike + swim

	run = run / total * 100
	bike = bike / total * 100
	swim = swim / total * 100

	return
}

// func secondToH(secondsIn float64) float64 {
// 	minutes := int(secondsIn) / 60
// 	seconds := int(secondsIn) % 60

// 	return 0.0
// }
