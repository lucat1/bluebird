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
var regText = "La #parola della #ghigliottina de #leredita di oggi è: (.*?)\n"
var reg, _ = regexp.Compile(regText)
var winnersRegText = ".+ @(.*?) - (.*?)\n.+ @(.*?) - (.*?)\n.+ @(.*?) - (.*?)($|\n)"
var winnersReg, _ = regexp.Compile(winnersRegText)

const timeFormat = "15:04:05"

func Ghigliottina(startTime, endTime *time.Time) (res GhigliottinaResponse, err error) {
	tweets, err := TweetsByUser("quizzettone", 50, startTime, endTime)
	if err != nil {
		return
	}
	var tweet Tweet
	found := false
	for i := len(tweets) - 1; i >= 0; i-- {
		t := tweets[i]
		if strings.Contains(t.Text, sub) {
			match := reg.FindStringSubmatch(t.Text)
			if len(match) > 0 && (found == false || (tweet.CreatedAt.After(t.CreatedAt)) && startTime.Before(t.CreatedAt) && endTime.After(t.CreatedAt)) {
				found = true
				tweet = t
			}
		}
	}
	if !found {
		return res, errors.New("No tweets were found")
	}
	match := reg.FindStringSubmatch(tweet.Text)
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
	winnersRaw := winnersReg.FindStringSubmatch(winTweet.Text)
	if len(winnersRaw) < 7 {
		return res, errors.New("Winners are malformed")
	}
	firstTime, errFirst := time.Parse(timeFormat, winnersRaw[2])
	secondTime, errSecond := time.Parse(timeFormat, winnersRaw[4])
	thirdTime, errThird := time.Parse(timeFormat, winnersRaw[6])
	if errFirst != nil || errSecond != nil || errThird != nil {
		return res, errors.New("Times are malformed")
	}
	res.Podium.First = GhigliottinaWinner{Username: winnersRaw[1], Time: firstTime}
	res.Podium.Second = GhigliottinaWinner{Username: winnersRaw[3], Time: secondTime}
	res.Podium.Third = GhigliottinaWinner{Username: winnersRaw[5], Time: thirdTime}

	return
}
