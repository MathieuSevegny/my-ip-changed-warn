package data

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type DataProvider struct {
	FilePath string
}

func (d *DataProvider) Save(content string) error {
	if err := d.EnsureCreated(); err != nil {
		return fmt.Errorf("error ensuring data file is created: %w", err)
	}

	if err := os.Truncate(d.FilePath, 0); err != nil {
		return fmt.Errorf("error truncating file: %w", err)
	}

	file, err := os.OpenFile(d.FilePath, os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	if _, err = io.WriteString(file, content); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func (d *DataProvider) Get() (string, error) {
	if err := d.EnsureCreated(); err != nil {
		return "", fmt.Errorf("error ensuring data file is created: %w", err)
	}

	file, err := os.Open(d.FilePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}

func (d *DataProvider) EnsureCreated() error {
	folder := filepath.Dir(d.FilePath)
	//If folder does not exist, create it
	if _, err := os.Stat(folder); os.IsNotExist(err) {
		if err = os.MkdirAll(folder, os.ModePerm); err != nil {
			return fmt.Errorf("error creating folder: %w", err)
		}
	}
	//If file does not exist, create it
	if _, err := os.Stat(d.FilePath); os.IsNotExist(err) {
		if _, err := os.Create(d.FilePath); err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}
	}
	return nil
}
