package src

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type CacheProvider struct {
	folder   string
	filename string
}

func (cache *CacheProvider) GetFilePath() string {
	return filepath.Join(cache.folder, cache.filename)
}

func (cache *CacheProvider) Save(content string) error {
	cache.EnsureCreated()
	filepath := cache.GetFilePath()
	err := os.Truncate(filepath, 0)
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(file, content)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (cache *CacheProvider) Get() (string, error) {
	cache.EnsureCreated()
	file_path := cache.GetFilePath()

	file, err := os.Open(file_path)
	if err != nil {
		log.Fatal(err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(content), nil
}

func (cache *CacheProvider) EnsureCreated() {
	//If folder does not exist, create it
	if _, err := os.Stat(cache.folder); os.IsNotExist(err) {
		err = os.MkdirAll(cache.folder, os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}
	}

	file_path := cache.GetFilePath()

	//If file does not exist, create it
	if _, err := os.Stat(file_path); os.IsNotExist(err) {
		_, err := os.Create(file_path)

		if err != nil {
			log.Fatal(err)
		}
	}
}
