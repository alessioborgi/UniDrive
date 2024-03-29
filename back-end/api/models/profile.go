package models

type Profile struct {
	Id                string  `json:"id"`
	Name              string  `json:"name"`
	Surname           string  `json:"surname"`
	Age               int     `json:"age"`
	Faculty           string  `json:"faculty"`
	University        string  `json:"university"`
	ProfilePictureUrl string  `json:"profile_picture_url"`
	Bio               string  `json:"bio"`
	Rating            float64 `json:"rating"`
	// Reviews []Review `json:"reviews"`
	TotalRidesGiven int        `json:"total_rides_given"`
	JoinedIn        string     `json:"joined_in"`
	CarDetails      CarDetails `json:"car_details"`
	City            string     `json:"city"`
	District        string     `json:"district"`
	InstagramName   string     `json:"instagram_name"`
	InstagramUrl    string     `json:"instagram_url"`
}

type Profile_DB struct {
	ID                string `gorm:"primary_key" json:"id"`
	Name              string `json:"name"`
	Surname           string `json:"surname"`
	Age               int    `json:"age"`
	Faculty           string `json:"faculty"`
	University        string `json:"university"`
	ProfilePictureUrl string `json:"profile_picture_url"`
	Bio               string `json:"bio"`
	JoinedIn          string `json:"joined_in"`
	City              string `json:"city"`
	District          string `json:"district"`
	InstagramName     string `json:"instagram_name"`
	InstagramUrl      string `json:"instagram_url"`
}
