/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random dad joke ",
	Long:  `This command fetches a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	url := "https://icanhazdadjoke.com/"
	bytes := getJokeData(url)
	joke := Joke{}

	err := json.Unmarshal(bytes, &joke)
	if err != nil {
		err = fmt.Errorf("unmarshaling from JSON to struct %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseApi string) []byte {
	request, err := http.NewRequest(http.MethodGet, baseApi, nil)
	if err != nil {
		err = fmt.Errorf("request a dadjoke: %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dadjoke CLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		err = fmt.Errorf("making a reaquest: %v", err)
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("reading a response body: %v", err)
	}

	return responseBytes
}
