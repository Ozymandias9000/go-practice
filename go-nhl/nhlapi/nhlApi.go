package nhlApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Team struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Link  string `json:"link"`
	Venue struct {
		Name     string `json:"name"`
		Link     string `json:"link"`
		City     string `json:"city"`
		TimeZone struct {
			ID     string `json:"id"`
			Offset int    `json:"offset"`
			Tz     string `json:"tz"`
		} `json:"timeZone"`
	} `json:"venue"`
	Abbreviation    string `json:"abbreviation"`
	TeamName        string `json:"teamName"`
	LocationName    string `json:"locationName"`
	FirstYearOfPlay string `json:"firstYearOfPlay"`
	Division        struct {
		ID           int    `json:"id"`
		Name         string `json:"name"`
		NameShort    string `json:"nameShort"`
		Link         string `json:"link"`
		Abbreviation string `json:"abbreviation"`
	} `json:"division"`
	Conference struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"conference"`
	Franchise struct {
		FranchiseID int    `json:"franchiseId"`
		TeamName    string `json:"teamName"`
		Link        string `json:"link"`
	} `json:"franchise"`
	ShortName       string `json:"shortName"`
	OfficialSiteURL string `json:"officialSiteUrl"`
	FranchiseID     int    `json:"franchiseId"`
	Active          bool   `json:"active"`
}

type nhlTeamsResponse struct {
	Teams []Team
}

type Player struct {
	Person struct {
		ID       int    `json:"id"`
		FullName string `json:"fullName"`
		Link     string `json:"link"`
	} `json:"person"`
	JerseyNumber string `json:"jerseyNumber"`
	Position     struct {
		Code         string `json:"code"`
		Name         string `json:"name"`
		Type         string `json:"type"`
		Abbreviation string `json:"abbreviation"`
	} `json:"position"`
}

type nhlRosterResponse struct {
	Roster []Player `json:"roster"`
}

const BaseURL = "https://statsapi.web.nhl.com/api/v1/"

func GetAllTeams() ([]Team, error) {
	res, err := http.Get(fmt.Sprintf("%s/teams", BaseURL))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response nhlTeamsResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Teams, err
}

func GetRoster(t Team) ([]Player, error) {
	res, err := http.Get(fmt.Sprintf("%s/teams/%v/roster", BaseURL, t.ID))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response nhlRosterResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Roster, err
}
