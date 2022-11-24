package request

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type GhigliottinaResponse struct {
	Word   string
	Podium GhigliottinaPodium
}

type GhigliottinaPodium struct {
	First  GhigliottinaWinner
	Second GhigliottinaWinner
	Third  GhigliottinaWinner
}

type GhigliottinaWinner struct {
	Username string
	Time     time.Time
}

var sub = "La #parola della #ghigliottina de #leredita di oggi è:"
var reg = "La #parola della #ghigliottina de #leredita di oggi è:(.*?)\n"
var winnersReg = ".+ @(.*?) - (.*?)\n.+ @(.*?) - (.*?)\n.+ @(.*?) - (.*?)($|\n)"

func Ghigliottina(startTime string, endTime string) (res GhigliottinaResponse, err error) {
	tweets, err := TweetsByUser("quizzettone", 50, startTime, endTime)
	if err != nil {
		return
	}
	start, err := time.Parse(time.RFC3339, startTime)
	if err != nil {
		return
	}
	end, err := time.Parse(time.RFC3339, endTime)
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
			if len(match) > 0 && (found == false || (tweet.CreatedAt.After(t.CreatedAt)) && start.Before(t.CreatedAt) && end.After(t.CreatedAt)) {
				found = true
				tweet = t
			}
		}
	}
	if found {
		match := r.FindStringSubmatch(tweet.Text)
		res.Word = (strings.Trim(match[1], " "))

		var tweetsReplies []Tweet
		tweetsReplies, err = Replies(tweet.ID, 50, "", "")
		if err != nil {
			return
		}
		if len(tweetsReplies) > 0 {
			winTweet := (tweetsReplies[len(tweetsReplies)-1])
			w, _ := regexp.Compile(winnersReg)
			winnersRaw := w.FindStringSubmatch(winTweet.Text)
			if len(winnersRaw) < 7 {
				return res, fmt.Errorf("Error while parsing winners data")
			}
			firstTime, err := time.Parse("15:04:05", winnersRaw[2])
			if err != nil {
				return res, err
			}
			secondTime, err := time.Parse("15:04:05", winnersRaw[4])
			if err != nil {
				return res, err
			}
			thirdTime, err := time.Parse("15:04:05", winnersRaw[6])
			if err != nil {
				return res, err
			}
			res.Podium.First = GhigliottinaWinner{Username: winnersRaw[1], Time: firstTime}
			res.Podium.Second = GhigliottinaWinner{Username: winnersRaw[3], Time: secondTime}
			res.Podium.Third = GhigliottinaWinner{Username: winnersRaw[5], Time: thirdTime}

			return res, nil
		}
	}
	return res, fmt.Errorf("Something went wrong")
}
