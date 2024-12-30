package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

func main() {
	//countDown(1000)
	eliteClubs("England", 20, 5)
}

type Club struct {
	Name                        string `json:"name"`
	Nation                      string `json:"nation"`
	EstimatedValueNumeric   int64 `json:"estimated_value_numeric"`
	NumberOfLeagueTitlesWon int64 `json:"number_of_league_titles_won"`
}

type ByValuationAndName []Club

func (m ByValuationAndName) Len() int { return len(m) }
func (m ByValuationAndName) Less(i, j int) bool {
	if m[i].EstimatedValueNumeric == m[j].EstimatedValueNumeric {
		return m[i].Name < m[j].Name
	}
	return m[i].EstimatedValueNumeric > m[j].EstimatedValueNumeric
}
func (m ByValuationAndName) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

type Response struct {
	Page       int64  `json:"page"`
	PerPage    int64  `json:"per_page"`
	TotalPages int64  `json:"total_pages"`
	Data       []Club `json:"data"`
}

func eliteClubs(nation string, minValuation int32, minTitlesWon int32) []string {
	var listOfClubs []string

	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/football_teams?nation=%s&estimated_value_numeric=%s&gt=%d&number_of_league_titles_won=%d&gt=%d", nation, minValuation, minValuation, minTitlesWon, minTitlesWon)
	response, err := http.Get(url)
	if err != nil {
		return listOfClubs
	}
	defer response.Body.Close()
	var resp Response
	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return listOfClubs
	}

	var clubs []Club
	if resp.TotalPages > resp.Page {
		clubs, err = retrieveClubs(url, resp.TotalPages, resp.Page+1, resp.Data)
	}

	sort.Sort(ByValuationAndName(clubs))

	clubNames := []string{}

	for _, club := range clubs {
		clubNames = append(clubNames, club.Name)
	}

	for i := 0; i < ; i++ {
		
	}

	return clubNames
}

func retrieveClubs(url string, totalPage int64, pageNum int64, clubs []Club) ([]Club, error) {
	if pageNum > totalPage {
		return clubs, nil
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	var resp Response

	err = json.NewDecoder(response.Body).Decode(&resp)

	if err != nil {
		return nil, err
	}

	return retrieveClubs(url, totalPage, pageNum+1, append(clubs, resp.Data...))
}
