package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type RaceInfo struct {
	Season    string    `json:"season"`
	Races     []Race    `json:"races"`
	Standings Standings `json:"standings"`
}

type Race struct {
	Round    int     `json:"round"`
	RaceName string  `json:"raceName"`
	Circuit  Circuit `json:"circuit"`
	Date     string  `json:"date"`
	Time     string  `json:"time"`
}

type Circuit struct {
	CircuitName string   `json:"circuitName"`
	Location    Location `json:"location"`
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
	Position   int    `json:"position"`
	DriverName string `json:"driverName"`
	Team       string `json:"team"`
	Points     int    `json:"points"`
}

type Constructors struct {
	Position int    `json:"position"`
	Team     string `json:"team"`
	Points   int    `json:"points"`
}

func ListRace() error {
	file, err := os.OpenFile("assets/mock.json", os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %v", err)
	}

	if fileStat.Size() <= 0 {
		fmt.Println("No race data")
		return nil
	}

	var RaceInfo []RaceInfo

	err = json.NewDecoder(file).Decode(&RaceInfo)
	if err != nil {
		return fmt.Errorf("failed to Decode file: %v", err)
	}

	
	return nil
}
