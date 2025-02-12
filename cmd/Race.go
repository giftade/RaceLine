package cmd 

type RaceInfo struct {
	Season    string    `json:"season"`
	Races     []Race    `json:"races"`
	Standings Standings `json:"standings"`
}

type Race struct {
	Round   int     `json:"round"`
	Name    string  `json:"name"`
	Circuit Circuit `json:"circuit"`
	Date    string  `json:"date"`
	Time    string  `json:"time"`
}

type Circuit struct {
	Name     string   `json:"name"`
	Location Location `json:"location"`
}

type Location struct {
	Country string `json:"country"`
	City    string `json:"city"`
}

type Standings struct {
	Drivers      []Drivers      `json:"drivers"`
	Constructors []Constructors `json:"constructors"`
}

type Drivers struct {
	Position int    `json:"position"`
	Name     string `json:"name"`
	Team     string `json:"team"`
	Points   int    `json:"points"`
}

type Constructors struct {
	Position int    `json:"position"`
	Team     string `json:"team"`
	Points   int    `json:"points"`
}
