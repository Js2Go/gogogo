package models

type Website struct {
	ID      int `json:"id" form:"id"`
	Name    string `json:"name" form:"name"`
	Url     string `json:"url" form:"url"`
	Alexa   string `json:"alexa" form:"alexa"`
	Country string `json:"country" form:"country"`
}
