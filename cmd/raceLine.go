package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

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

func schedule() error {
	file, err := openFile("assets/mock.json")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	var RaceInfo []RaceInfo

	err = json.NewDecoder(file).Decode(&RaceInfo)
	if err != nil {
		fmt.Printf("failed to Decode file: %v", err)
		return err
	}
	printRaceSchedule(RaceInfo)
	return nil
}

func printRaceSchedule(RaceInfo []RaceInfo) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Round", "Race Name", "Circuit", "Country", "City", "Date", "Time"})

	for _, season := range RaceInfo {
		fmt.Println("Season:", season.Season)

		for _, race := range season.Races {
			table.Append([]string{
				fmt.Sprintf("%d", race.Round),
				race.Name,
				race.Circuit.Name,
				race.Circuit.Location.Country,
				race.Circuit.Location.City,
				race.Date,
				race.Time,
			})
		}

		table.Render() // Print the table
		fmt.Println()
	}
}
