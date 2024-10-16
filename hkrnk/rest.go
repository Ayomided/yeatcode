package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Matches struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	Total      int         `json:"total"`
	TotalPages int         `json:"total_pages"`
	Data       []MatchData `json:"data"`
}

type MatchData struct {
	Competition string `json:"competition"`
	Year        int    `json:"year"`
	Round       string `json:"round"`
	Team1       string `json:"team1"`
	Team2       string `json:"team2"`
	Team1Goals  string `json:"team1goals"`
	Team2Goals  string `json:"team2goals"`
}

type Competitions struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Data       []struct {
		Name     string `json:"name"`
		Country  string `json:"country"`
		Year     int    `json:"year"`
		Winner   string `json:"winner"`
		Runnerup string `json:"runnerup"`
	} `json:"data"`
}

const baseURL = "https://jsonmock.hackerrank.com/api/"

func apiCaller(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return io.ReadAll(response.Body)
}

func GetNumDraws(year int) (int, error) {
	start := time.Now()
	var totalDraws int

	apiURL := fmt.Sprintf("%sfootball_matches?year=%d", baseURL, year)

	for page := 1; ; page++ {
		url := fmt.Sprintf("%s&page=%d", apiURL, page)
		responseBody, err := apiCaller(url)
		if err != nil {
			return 0, err
		}

		var matches Matches
		if err := json.Unmarshal(responseBody, &matches); err != nil {
			return 0, err
		}

		for _, match := range matches.Data {
			if match.Team1Goals == match.Team2Goals {
				totalDraws++
			}
		}

		if page >= matches.TotalPages {
			break
		}
	}

	fmt.Printf("Total draws in %d: %d\n", year, totalDraws)
	fmt.Printf("Time taken: %v\n", time.Since(start))

	return totalDraws, nil
}

func GetWinnerTotalGoals(competition string, year int) (int, error) {
	start := time.Now()

	// Get competition details
	compURL := fmt.Sprintf("%sfootball_competitions?name=%s&year=%d", baseURL, url.QueryEscape(competition), year)
	response, err := apiCaller(compURL)
	if err != nil {
		return 0, err
	}

	var competitions Competitions
	if err := json.Unmarshal(response, &competitions); err != nil {
		return 0, err
	}

	if len(competitions.Data) == 0 {
		return 0, fmt.Errorf("no competition found for %s in %d", competition, year)
	}

	winner := competitions.Data[0].Winner
	totalGoals := 0

	// Function to get goals for a team (as team1 or team2)
	getGoals := func(team string) (int, error) {
		goals := 0
		for page := 1; ; page++ {
			url := fmt.Sprintf("%sfootball_matches?competition=%s&year=%d&team%s=%s&page=%d",
				baseURL, url.QueryEscape(competition), year, team, url.QueryEscape(winner), page)

			response, err := apiCaller(url)
			if err != nil {
				return 0, err
			}

			var matches Matches
			if err := json.Unmarshal(response, &matches); err != nil {
				return 0, err
			}

			for _, match := range matches.Data {
				if team == "1" {
					goals += parseInt(match.Team1Goals)
				} else {
					goals += parseInt(match.Team2Goals)
				}
			}

			if page >= matches.TotalPages {
				break
			}
		}
		return goals, nil
	}

	// Get goals as team1
	goalsAsTeam1, err := getGoals("1")
	if err != nil {
		return 0, err
	}
	totalGoals += goalsAsTeam1

	// Get goals as team2
	goalsAsTeam2, err := getGoals("2")
	if err != nil {
		return 0, err
	}
	totalGoals += goalsAsTeam2

	fmt.Printf("Total goals for %s in %s %d: %d\n", winner, competition, year, totalGoals)
	fmt.Printf("Time taken: %v\n", time.Since(start))

	return totalGoals, nil
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
