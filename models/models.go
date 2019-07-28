package models

type Configuration struct {
	Token        string
	ClientId     string
	ClientSecret string
	RefreshToken string
}

type RefreshedToken struct {
	Token_Type    string
	Access_Token  string
	Expires_At    int
	Expires_In    int
	Refresh_Token string
}

type Response struct {
	Message string
}

type Activities struct {
	Act []Activity
}

type Activity struct {
	Name        string
	Distance    float64
	Moving_Time int
	Type        string
}

type SingleActivity struct {
	Split_Metrics []SplitMetrics `json:"splits_metrics"`
	Laps          []Laps         `json:"laps"`
}

type SplitMetrics struct {
	Distance             float64 `json:"distance"`
	Elapsed_Time         int     `json:"elapsed_time"`
	Elevation_Difference float64 `json:"elevation_difference"`
	Moving_Time          int     `json:"moving_time"`
	Split                int     `json:"split"`
	Average_Speed        float64 `json:"average_speed"`
}

type Laps struct {
	Distance      float64 `json:"distance"`
	Elapsed_Time  int     `json:"elapsed_time"`
	Moving_Time   int     `json:"moving_time"`
	Split         int     `json:"split"`
	Average_Speed float64 `json:"average_speed"`
	Max_Speed     float64 `json:"max_speed"`
	Average_Watts float64 `json:"average_watts"`
}

type ActivityType int

const (
	Run  ActivityType = 0
	Ride ActivityType = 1
	Swim ActivityType = 2
)
