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
	startTimeStr := ""
	if err == nil {
		startTimeStr = startTime.Format(time.RFC3339)
	}
	endTime, err := time.Parse(time.RFC3339, rawEndTime)
	endTimeStr := ""
	if err == nil {
		endTimeStr = endTime.Format(time.RFC3339)
	}
	res, err := request.Ghigliottina(startTimeStr, endTimeStr)
	if err != nil {
		sendError(w, 500, APIError{
			Message: "Could not process ghigliottina game",
			Error:   err,
		})
		return
	}
	sendJSON(w, http.StatusOK, res)
}
