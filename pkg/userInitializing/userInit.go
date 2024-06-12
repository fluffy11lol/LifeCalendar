package userInitializing

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name      string `json:"name"`
	Birthdate struct {
		Day   int `json:"day"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"birthdate"`
}

func (u *User) InitInfoFromJSON() error {
	jsonFile, err := os.ReadFile("../userInfo.json")
	if err != nil {
		jsonFile, err = os.ReadFile("userInfo.json")
	}
	if err != nil {
		return fmt.Errorf("Error reading JSON file: %v\n\n", err)
	}
	err = json.Unmarshal(jsonFile, &u)
	if err != nil {
		return fmt.Errorf("Error deserialization JSON: %v\n", err)
	}
	return nil
}
