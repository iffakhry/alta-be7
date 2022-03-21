package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StarWars struct {
	Result []StarWarsPeople `json:"results"`
}

type StarWarsPeople struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	SkinColor string   `json:"skin_color"`
	BirthYear string   `json:"birth_year"`
	Films     []string `json:"films"`
}

func main() {
	response, err := http.Get("https://swapi.dev/api/people")
	if err != nil {
		fmt.Println("error to get data")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("failed to read data")
	}
	defer response.Body.Close()

	var peoples StarWars
	// var people1 StarWarsPeople
	errJSON := json.Unmarshal(responseData, &peoples)
	if errJSON != nil {
		fmt.Println("Error unmarshal", errJSON)
	}

	// fmt.Println("name", people1.Name)
	fmt.Println("name", peoples.Result[3].Name)
	fmt.Println("birth year", peoples.Result[3].BirthYear)
	fmt.Println("skin color", peoples.Result[3].SkinColor)
}
