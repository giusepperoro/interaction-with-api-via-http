package entity

type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

type Street struct {
	Number int64  `json:"number"`
	Name   string `json:"name"`
}

type Location struct {
	Gender string `json:"gender"`
	Street Street `json:"street"`
}

type Coordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Result struct {
	Gender      string      `json:"gender"`
	Name        Name        `json:"name"`
	Location    Location    `json:"location"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Postcode    string      `json:"postcode"`
	Coordinates Coordinates `json:"coordinates"`
}
