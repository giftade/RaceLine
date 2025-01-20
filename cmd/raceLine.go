package cmd

import (
	"fmt"
	"os"
)

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
		fmt.Println("No expenses to list")
		return nil
	}

	return nil
}
