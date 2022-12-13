package server

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"git.hjkl.gq/team14/team14/cache"
	"git.hjkl.gq/team14/team14/request"
)

type PoliticiansScoreResponse struct {
	Politicians []request.Politician `json:"politicians"`
}

type PoliticiansScoreboardResponse struct {
	Politicians     []request.Politician `json:"politicians"`
	BestClimber     request.Politician   `json:"best_climber"`
	BestAverage     request.Politician   `json:"best_average"`
	BestSingleScore request.Politician   `json:"best_single_score"`
}

type TeamResponse struct {
	Username   string   `json:"username"`
	PictureURL string   `json:"picture_url"`
	Name       string   `json:"name,omitempty"`
	Leader     string   `json:"leader,omitempty"`
	Members    []string `json:"members"`
}

type TeamsResponse struct {
	Teams []request.Team `json:"teams"`
}

func politiciansScoreHandler(w http.ResponseWriter, r *http.Request) {
	politiciansScore := []request.Politician{}
	var err error
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		sendError(w, http.StatusBadRequest, APIError{
			Message: "Invalid amount query",
			Error:   err,
		})
		return
	}

	rawStartTime := r.URL.Query().Get("startTime")
	rawEndTime := r.URL.Query().Get("endTime")

	startTime, err := time.Parse(time.RFC3339, rawStartTime)
	if err != nil {
		sendError(w, http.StatusBadRequest, APIError{
			Message: "Start time is wrong",
			Error:   err,
		})
		return
	}
	endTime, err := time.Parse(time.RFC3339, rawEndTime)
	if err != nil {
		sendError(w, http.StatusBadRequest, APIError{
			Message: "End time is wrong",
			Error:   err,
		})
		return
	}
	politiciansScore, err = request.PoliticiansScore(uint(amount), startTime, endTime)
	sendJSON(w, http.StatusOK, PoliticiansScoreResponse{
		Politicians: politiciansScore,
	})
}

func politiciansScoreboardHandler(w http.ResponseWriter, r *http.Request) {
	politiciansOrdered, err := cache.PoliticiansScoreboard()
	if err != nil {
		sendError(w, http.StatusBadRequest, APIError{
			Message: "Cannot fetch politicians scoreboard from database",
			Error:   err,
		})
		return
	}
	bestAverage, err := cache.PoliticianBestAverage()
	if err != nil {
		sendError(w, http.StatusBadRequest, APIError{
			Message: "Cannot fetch best avereage politician from database",
			Error:   err,
		})
		return
	}
	bestSingleScore, err := cache.PoliticianBestSingleScore()
	if err != nil {
		sendError(w, http.StatusBadRequest, APIError{
			Message: "Cannot fetch best avereage politician from database",
			Error:   err,
		})
		return
	}
	sendJSON(w, http.StatusOK, PoliticiansScoreboardResponse{
		Politicians:     politiciansOrdered,
		BestAverage:     bestAverage,
		BestSingleScore: bestSingleScore,
	})
}

func searchHandlerTeam(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username != "" {
		team, err := cache.SearchTeamByUsername(username)
		if err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Team not found",
				Error:   err,
			})
			return
		}

		var teamInfo request.OCRTeam
		if ocrEnable, err := strconv.ParseBool(r.URL.Query().Get("ocr")); err == nil && ocrEnable == true {
			teamInfo, _ = request.GetTeamInfo(team.PictureURL)
		}

		sendJSON(w, http.StatusOK, TeamResponse{Username: team.Username, PictureURL: team.PictureURL, Name: teamInfo.Name, Leader: teamInfo.Leader, Members: teamInfo.Members})
		return
	}

	sendError(w, http.StatusBadRequest, APIError{
		Message: "Username is mandatory",
		Error:   fmt.Errorf(""),
	})
}

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := cache.TeamsAll()
	if err != nil {
		sendError(w, http.StatusBadRequest, APIError{
			Message: "Teams not found",
			Error:   err,
		})
		return
	}

	sendJSON(w, http.StatusOK, TeamsResponse{Teams: teams})
}
