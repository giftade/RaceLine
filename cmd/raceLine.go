package cmd

import (
	"encoding/json"
	"fmt"
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
