package cmd

import (
	"errors"
	"os"
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
