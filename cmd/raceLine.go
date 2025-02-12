package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

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
