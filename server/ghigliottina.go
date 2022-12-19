package server

import (
	"net/http"
	"time"

	"git.hjkl.gq/team14/team14/request"
)

func getGhigliottina(w http.ResponseWriter, r *http.Request) {
	rawStartTime := r.URL.Query().Get("startTime")
	rawEndTime := r.URL.Query().Get("endTime")
	endTime, err := time.Parse(time.RFC3339, rawEndTime)
	if err != nil {
		endTime = time.Now().Add(time.Minute * time.Duration(-1))
	}
	startTime, err := time.Parse(time.RFC3339, rawStartTime)
	if err != nil {
		startTime = time.Now().AddDate(0, 0, -1)
	}
	res, err := request.Ghigliottina(&startTime, &endTime)
	if err != nil {
		sendError(w, 500, APIError{
			Message: "Could not process ghigliottina game",
			Error:   err,
		})
		return
	}
	sendJSON(w, http.StatusOK, res)
}
