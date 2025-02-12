package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func openFile(addr string) (*os.File, error) {
	file, err := os.OpenFile(addr, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	if fileStat.Size() <= 0 {
		return nil, errors.New("no race data")
	}

	return file, nil

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
