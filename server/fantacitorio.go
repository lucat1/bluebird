package server

import (
	"net/http"
	"strconv"

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

func politiciansScoreHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	politiciansScore := []request.Politician{}
	if query != "" {
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

		politiciansScore, err = request.PoliticiansScore(uint(amount), rawStartTime, rawEndTime)
	}
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
