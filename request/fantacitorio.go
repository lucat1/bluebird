package request

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const MAXPOLITICIANS = 761
const FANTAUSER = "Fanta_citorio"
const TEAMSTARTTIME = "2022-10-27T22:00:20.602Z"
const TEAMENDTIME = "2022-10-28T22:00:20.602Z"

var startTime, _ = time.Parse(time.RFC3339, TEAMSTARTTIME)
var endTime, _ = time.Parse(time.RFC3339, TEAMENDTIME)

const regPuntiTxt = "([0-9O]*?) PUNTI - (.*?)$|([0-9O]*?) PUNTI PER (.*?)$|([0-9O]*?) PUNTI A (.*?)$|.*? TOT\\.([0-9O]*?) - (.*?)$|ALTRI ([0-9O]*?) PUNTI PER (.*?)$"
const regPuntiRevTxt = "(.*?) ([0-9O]*?) punti$"

var regPunti, _ = regexp.Compile(regPuntiTxt)
var regPuntiRev, _ = regexp.Compile(regPuntiRevTxt)

func searchNameBySurname(surname string) (name string, err error) {
	for i := 0; i < MAXPOLITICIANS; i++ {
		if strings.ToUpper(allPoliticians[i][1]) == surname {
			return strings.ToUpper(allPoliticians[i][0]), nil
		}
	}
	return "", fmt.Errorf("Name not found")
}

func checkNameSurname(nameSurname string) (name string, surname string, err error) {
	split := strings.Split(nameSurname, " ")
	if len(split) < 2 {
		// in case of "" as nameSurname split[0] will be ""
		surname = split[0]
		name, err = searchNameBySurname(surname)
		if err != nil {
			return
		}
	} else {
		name = split[0]
		surname = split[1]
		var nameCheck string
		nameCheck, err = searchNameBySurname(surname)
		if err != nil {
			return
		}
		if name != nameCheck {
			return "", "", fmt.Errorf("Invalid name and surname")
		}
	}

	return
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func parsePolitician(text string) (res Politician, err error) {
	nameSurname := ""
	points := 0
	match := removeEmptyStrings(regPunti.FindStringSubmatch(text))
	if len(match) >= 2 {
		nameSurname = match[2]
		// cant cause an error due to how regexp is defined
		points, _ = strconv.Atoi(strings.ReplaceAll(match[1], "O", "0"))
	} else {
		match := removeEmptyStrings(regPuntiRev.FindStringSubmatch(text))
		if len(match) >= 2 {
			nameSurname = match[1]
			// cant cause an error due to how regexp is defined
			points, _ = strconv.Atoi(strings.ReplaceAll(match[2], "O", "0"))
		}
	}

	if nameSurname != "" {
		var name string
		var surname string
		name, surname, err = checkNameSurname(nameSurname)
		if err != nil {
			return
		}
		res = Politician{ID: 0, Name: name, Surname: surname, Points: points, BestSingleScore: points, Average: float64(points), NPosts: 1}
	} else {
		return res, fmt.Errorf("Error while parsing politician")
	}
	return
}

func AddPoliticianToList(politician Politician, list []Politician) []Politician {

	found := false
	for i := range list {
		if list[i].Name == politician.Name && list[i].Surname == politician.Surname {
			list[i].Points += politician.Points
			list[i].NPosts += 1
			if list[i].BestSingleScore < politician.Points {
				list[i].BestSingleScore = politician.Points
			}
			list[i].Average = float64(list[i].Points) / float64(list[i].NPosts)
			found = true
		}
	}
	if !found {
		list = append(list, politician)
	}

	return list
}

// week by week , so 4/5 posts per day -> 30/40 posts per week
func PoliticiansScore(n uint, startTime time.Time, endTime time.Time) (politicians []Politician, err error) {
	tweets, err := TweetsByUser(FANTAUSER, n, &startTime, &endTime)
	if err != nil {
		return
	}
	for _, t := range tweets {
		split := strings.Split(t.Text, "\n")
		for _, s := range split {
			politician, err := parsePolitician(s)
			if err != nil {
				continue
			}
			politician.LastUpdated = t.CreatedAt
			politicians = AddPoliticianToList(politician, politicians)
		}
	}
	return politicians, nil
}

func Teams() (teams []Team, err error) {
	// cant go wrong
	tweets, err := TweetsByUser(FANTAUSER, 500, &startTime, &endTime)
	if err != nil {
		return
	}
	for _, t := range tweets {
		if len(*t.Media) > 0 && len(*t.Mentions) > 0 {
			teams = append(teams, Team{Username: (*t.Mentions)[0].Username, PictureURL: (*t.Media)[0].URL})
		}
	}
	return
}
