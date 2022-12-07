package server

import (
	"git.hjkl.gq/team14/team14/request"
	"github.com/kataras/muxie"
	"log"
	"net/http"
)

const (
	nOfAPITweets   uint = 100
	nOfDaysAllowed      = 7
)

type indexPayload struct {
	Query  string
	Tweets []request.Tweet
}

func cors(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		f.ServeHTTP(w, r)
	})
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
}

func RunServer(host string) error {
	mux := muxie.NewMux()

	mux.Use(cors)
	mux.HandleFunc("/api/search", searchHandler)
	mux.HandleFunc("/api/sentiment", sentimentHandler)
	mux.HandleFunc("/api/chess", chessHandler)
	mux.HandleFunc("/api/ghigliottina", getGhigliottina)
	mux.HandleFunc("/api/fantacitorio/scores", politiciansScoreHandler)
	mux.HandleFunc("/api/fantacitorio/classifica", politiciansScoreboardHandler)
	mux.Handle("/assets/*path", http.FileServer(http.Dir("dist")))
	mux.HandleFunc("/*path", serveIndex)

	log.Printf("Listening on %s\n", host)
	return http.ListenAndServe(host, mux)
}
