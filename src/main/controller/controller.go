package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/ank1106/webserver/src/main/models"
)

type Controllers struct {
	DB *models.DBClient
}

func NewController(db *models.DBClient) *Controllers {
	return &Controllers{
		DB: db,
	}
}

type Request struct {
	Type     string `json:"type"`
	Query    string `json:"query"`
	Language string `json:"language"`
	Unit     string `json:"unit"`
}

type Location struct {
	Name           string `json:"name"`
	Country        string `json:"country"`
	Region         string `json:"region"`
	Lat            string `json:"lat"`
	Lon            string `json:"lon"`
	TimezoneID     string `json:"timezone_id"`
	Localtime      string `json:"localtime"`
	LocaltimeEpoch int    `json:"localtime_epoch"`
	UtcOffset      string `json:"utc_offset"`
}

type Current struct {
	ObservationTime     string   `json:"observation_time"`
	Temperature         int      `json:"temperature"`
	WeatherCode         int      `json:"weather_code"`
	WeatherIcons        []string `json:"weather_icons"`
	WeatherDescriptions []string `json:"weather_descriptions"`
	WindSpeed           int      `json:"wind_speed"`
	WindDegree          int      `json:"wind_degree"`
	WindDir             string   `json:"wind_dir"`
	Pressure            int      `json:"pressure"`
	Precip              int      `json:"precip"`
	Humidity            int      `json:"humidity"`
	Cloudcover          int      `json:"cloudcover"`
	Feelslike           int      `json:"feelslike"`
	UvIndex             int      `json:"uv_index"`
	Visibility          int      `json:"visibility"`
	IsDay               string   `json:"is_day"`
}

type Weather struct {
	Request  Request  `json:"request"`
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

func (c *Controllers) IndexHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("<h1>Hello World!</h1>"))
	var tpl = template.Must(template.ParseFiles("src/main/templates/dash.html"))
	if r.Method == http.MethodGet {
		// cookie, err := r.Cookie("test")
		// if err != nil {
		// 	fmt.Printf("Cant find cookie :/\r\n")
		// } else {
		// 	fmt.Println(cookie)
		// }
		apiKey := "86490556399f5194031fd81d53004ccb"
		endpoint := fmt.Sprintf("http://api.weatherstack.com/current/?access_key=%s&query=india", apiKey)
		resp, err := http.Get(endpoint)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		defer resp.Body.Close()
		weather := Weather{}
		err = json.NewDecoder(resp.Body).Decode(&weather)
		fmt.Println(weather, resp.Body, err, resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = tpl.Execute(w, weather)
		if err != nil {
			log.Println(err)
		}

	}

	// tpl.Execute(w, nil)

}

func (c *Controllers) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("src/main/templates/login.html"))
	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			fmt.Println(w, "ParseForm() err: %v", err)
			return
		}
		email := r.PostForm["email"]
		password := r.PostForm["password"]

		if user := c.DB.AuthenticateUser(strings.Join(email, " "), strings.Join(password, " ")); user != nil {
			fmt.Println("successfully loggedin")
		}
	}
	tpl.Execute(w, nil)
}

func (c *Controllers) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var tpl = template.Must(template.ParseFiles("src/main/templates/register.html"))
	if r.Method == http.MethodPost {

		if err := r.ParseForm(); err != nil {
			fmt.Println(w, "ParseForm() err: %v", err)
			return
		}
		user := models.User{
			Email:    strings.Join(r.PostForm["email"], " "),
			Password: strings.Join(r.PostForm["password"], " "),
			Phone:    strings.Join(r.PostForm["phone"], " "),
			Name:     strings.Join(r.PostForm["name"], " "),
		}
		if err := c.DB.CreateUser(&user); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("User created")

	}
	tpl.Execute(w, nil)
}
