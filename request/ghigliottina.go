package request

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type GhigliottinaResponse struct {
	Word   string             `json:"word"`
	Podium GhigliottinaPodium `json:"podium"`
}

type GhigliottinaPodium struct {
	First  GhigliottinaWinner `json:"first"`
	Second GhigliottinaWinner `json:"second"`
	Third  GhigliottinaWinner `json:"third"`
}

type GhigliottinaWinner struct {
	Username string    `json:"username"`
	Time     time.Time `json:"time"`
}

var sub = "La #parola della #ghigliottina de #leredita di oggi è:"
var reg = "La #parola della #ghigliottina de #leredita di oggi è:(.*?)\n"
var winnersReg = ".+ @(.*?) - (.*?)\n.+ @(.*?) - (.*?)\n.+ @(.*?) - (.*?)($|\n)"

const timeFormat = "15:04:05"

func Ghigliottina(startTime, endTime *time.Time) (res GhigliottinaResponse, err error) {
	tweets, err := TweetsByUser("quizzettone", 50, startTime, endTime)
	if err != nil {
		return
	}
	r, _ := regexp.Compile(reg)
	var tweet Tweet
	found := false
	for i := len(tweets) - 1; i >= 0; i-- {
		t := tweets[i]
		if strings.Contains(t.Text, sub) {
			match := r.FindStringSubmatch(t.Text)
			if len(match) > 0 && (found == false || (tweet.CreatedAt.After(t.CreatedAt)) && startTime.Before(t.CreatedAt) && endTime.After(t.CreatedAt)) {
				found = true
				tweet = t
			}
		}
	}
	if !found {
		return res, errors.New("No tweets were found")
	}
	match := r.FindStringSubmatch(tweet.Text)
	res.Word = (strings.Trim(match[1], " "))

	var tweetsReplies []Tweet
	tweetsReplies, err = Replies(tweet.ID, 50, nil, nil)
	if err != nil {
		return
	}
	if len(tweetsReplies) <= 0 {
		return res, errors.New("No tweet replies were found")
	}

	winTweet := (tweetsReplies[len(tweetsReplies)-1])
	w, _ := regexp.Compile(winnersReg)
	winnersRaw := w.FindStringSubmatch(winTweet.Text)
	if len(winnersRaw) < 7 {
		return res, errors.New("Error while parsing winners data")
	}
	firstTime, err := time.Parse(timeFormat, winnersRaw[2])
	if err != nil {
		return
	}
	secondTime, err := time.Parse(timeFormat, winnersRaw[4])
	if err != nil {
		return
	}
	thirdTime, err := time.Parse(timeFormat, winnersRaw[6])
	if err != nil {
		return
	}
	res.Podium.First = GhigliottinaWinner{Username: winnersRaw[1], Time: firstTime}
	res.Podium.Second = GhigliottinaWinner{Username: winnersRaw[3], Time: secondTime}
	res.Podium.Third = GhigliottinaWinner{Username: winnersRaw[5], Time: thirdTime}

	return
}
