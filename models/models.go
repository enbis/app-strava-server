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

type ActivityType int

const (
	Run  ActivityType = 0
	Ride ActivityType = 1
	Swim ActivityType = 2
)
