package server

import (
	"git.hjkl.gq/team14/team14/request"
	"net/http"
	"time"
)

func getGhigliottina(w http.ResponseWriter, r *http.Request) {
	rawStartTime := r.URL.Query().Get("startTime")
	rawEndTime := r.URL.Query().Get("endTime")
	startTime, err := time.Parse(time.RFC3339, rawStartTime)
	if err == nil {
		startTime = time.Now()
	}
	endTime, err := time.Parse(time.RFC3339, rawEndTime)
	if err == nil {
		endTime = time.Now().Add(1)
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
