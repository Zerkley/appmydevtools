package helpers

import (
	"encoding/json"
	"fmt"
	"os"
)

type Database struct {
	Dev   string `json:"dev"`
	Stage string `json:"stage"`
}

type Settings struct {
	Text     string   `json:"text"`
	Database Database `json:"database"`
}

func createFile(path string) {
	fmt.Println("Creating a new file...")
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	settings := Settings{
		Text: "hola",
	}

	// Marshal the struct to JSON
	newJSON, err := json.MarshalIndent(settings, "", "  ")
	if err != nil {
		fmt.Printf("error marshaling JSON: %s", err)
		return
	}

	// Write the JSON data
	_, err = file.Write(newJSON)
	if err != nil {
		fmt.Printf("error writing to file: %s", err)
		return
	}

	fmt.Println("New settings file created and saved successfully.")

}

func SettingsLoader() (Settings, error) {
	path := "./tmp/settings.json"

	//Check if file exists
	_, err := os.Stat(path)

	// If it doesn't exist, create it with defaults
	if err != nil {
		fmt.Println("File doesn't exist")
		createFile(path)
	}

	// Open and read file
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening %s: %s", path, err)
		return Settings{}, err
	}
	defer file.Close()

	var settings Settings
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&settings)
	if err != nil {
		fmt.Printf("Error decoding JSON: %v", err)
		return Settings{}, err
	}

	return settings, nil
}
