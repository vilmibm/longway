package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Platform struct {
	ID   int    `json:"platform_id"`
	Name string `json:"platform_name"`
}

type gameData struct {
	Games []struct {
		Title     string
		Platforms []Platform
	}
}

func _main(apiKey string) error {
	var foundGame string
	var foundPlatform string

	v := url.Values{}
	v.Add("format", "normal")
	v.Add("api_key", apiKey)

	u := fmt.Sprintf("https://api.mobygames.com/v1/games/random?%s", v.Encode())

	res, err := http.Get(u)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return fmt.Errorf("sad satus: %d", res.StatusCode)
	}

	var gd gameData
	if err = json.NewDecoder(res.Body).Decode(&gd); err != nil {
		return err
	}

	rand.Seed(time.Now().Unix())

	choice := rand.Intn(len(gd.Games))

	for ix, game := range gd.Games {
		if ix != choice {
			continue
		}
		foundGame = game.Title
		foundPlatform = game.Platforms[0].Name
	}

	utterances := []string{
		"huh",
		"wow",
		"whoa",
		"well",
		"geez",
		"neato",
	}

	pauses := []string{
		"...",
		", ",
		". ",
		"! ",
		"; ",
	}

	suffices := []string{
		"",
		".",
		"!",
		"...",
	}

	utterance := utterances[rand.Intn(len(utterances))]
	pause := pauses[rand.Intn(len(pauses))]
	suffix := suffices[rand.Intn(len(suffices))]

	fmt.Printf("%s%sgames sure have come a long way since %s for %s%s\n", utterance, pause, foundGame, foundPlatform, suffix)

	return nil
}

func main() {
	apiKey := os.Getenv("MOBYKEY")
	if apiKey == "" {
		panic("MOBYKEY not found in env")
	}
	if err := _main(apiKey); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
