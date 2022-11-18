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
	username string
	Time     time.Time
}

var sub = "La #parola della #ghigliottina de #leredita di oggi è:"
var reg = "La #parola della #ghigliottina de #leredita di oggi è:(.*?)\n"
var winnersReg = ".+ @(.*?) - (.*?)\n.+ @(.*?) - (.*?)\n.+ @(.*?) - (.*?)[\n|$]"

func Ghigliottina() (res GhigliottinaResponse, err error) {
	tweets, err := TweetsByUser("quizzettone", 50, "", "")
	if err != nil {
		return
	}
	r, _ := regexp.Compile(reg)
	for _, t := range tweets {
		if strings.Contains(t.Text, sub) {
			match := r.FindStringSubmatch(t.Text)
			if len(match) > 0 {
				res.Word = (strings.Trim(match[1], " "))

				var tweetsReplies []Tweet
				tweetsReplies, err = Replies(t.ID, 50, "", "")
				if err != nil {
					return
				}
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
				res.Podium.First = GhigliottinaWinner{username: winnersRaw[1], Time: firstTime}
				res.Podium.Second = GhigliottinaWinner{username: winnersRaw[3], Time: secondTime}
				res.Podium.Third = GhigliottinaWinner{username: winnersRaw[5], Time: thirdTime}

				return res, nil
			}
		}
	}
	return
}
