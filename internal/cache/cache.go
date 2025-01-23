package cache

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type CacheProvider struct {
	Folder   string
	Filename string
}

func (cache *CacheProvider) GetFilePath() string {
	return filepath.Join(cache.Folder, cache.Filename)
}

func (cache *CacheProvider) Save(content string) error {
	cache.EnsureCreated()
	filepath := cache.GetFilePath()
	err := os.Truncate(filepath, 0)

	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)

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
	if cache.Folder == "" || cache.Filename == "" {
		cache.Folder = "./cache"
		cache.Filename = "last_ip.txt"
	}
	//If folder does not exist, create it
	if _, err := os.Stat(cache.Folder); os.IsNotExist(err) {
		err = os.MkdirAll(cache.Folder, os.ModePerm)

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
