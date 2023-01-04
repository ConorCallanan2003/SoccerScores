package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Invalid Request\n")
}
func getLeagues(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Printf("got /leagues request\n")
	io.WriteString(w, data)
}

var data string

func main() {

	data = string(mapToJSON(apiCaller()))

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/leagues", getLeagues)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func mapToJSON(leagues map[string][]Team) (result []byte) {
	var resultList []League
	for key, element := range leagues {
		var league League
		league.Name = key
		league.Teams = element
		resultList = append(resultList, league)
	}
	resultJSON, err := json.Marshal(resultList)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(resultJSON))

	return resultJSON
}

func apiCaller() (leagues map[string][]Team) {

	result := make(map[string][]Team)

	resp, err := http.Get("https://api.soccersapi.com/v2.2/leagues/?user=conor.callanan1&token=54fa83edd3897eccee4f88463a3e6a9e&t=list")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response Response
	json.Unmarshal([]byte(body), &response)

	for i := 0; i < len(response.Data); i++ {

		var league = response.Data[i]

		fmt.Printf("\n\nLeague Name: %s, League ID: %s, League Country: %s, Current Season ID: %s Requests Remaining: %d \n\n", league.Name, league.Id, league.Country_Name, league.Current_Season_Id, response.Meta.Requests_Left)

		resp, err := http.Get("https://api.soccersapi.com/v2.2/leagues/?user=conor.callanan1&token=54fa83edd3897eccee4f88463a3e6a9e&t=standings&season_id=" + league.Current_Season_Id)
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var response StandingsOuterResponse
		json.Unmarshal([]byte(body), &response)

		var teams = response.Data.Standings

		result[league.Name] = teams

		for i := 0; i < len(teams); i++ {
			fmt.Printf("Team Name: %s, Team Position: %d \n", teams[i].Team_Name, teams[i].Overall.Position)
		}

	}

	return result

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type Response struct {
	Data []League
	Meta Meta
}

type StandingsOuterResponse struct {
	Data StandingsResponse
	Meta Meta
}

type StandingsResponse struct {
	Standings []Team
}

type League struct {
	Id                string
	Current_Season_Id string
	Name              string
	Country_Name      string
	Teams             []Team
}

type Team struct {
	Team_Name string
	Overall   TeamStats
}

type TeamStats struct {
	Position     int
	Won          int
	Draw         int
	Lost         int
	Points       int
	Games_Played int
}

type Meta struct {
	Requests_Left int
	Pages         int
	Count         int
}
