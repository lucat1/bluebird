package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetTeamInfo(imageURL string) (team OCRTeam, err error) {
	client := &http.Client{}
	form := url.Values{}
	form.Add("url", imageURL)
	form.Add("language", "ita")
	form.Add("isOverlayRequired", "true")
	form.Add("FileType", ".Auto")
	form.Add("IsCreateSearchablePDF", "false")
	form.Add("isSearchablePdfHideTextLayer", "true")
	form.Add("detectOrientation", "false")
	form.Add("isTable", "false")
	form.Add("scale", "true")
	form.Add("OCREngine", "2")
	form.Add("detectCheckbox", "false")
	form.Add("checkboxTemplate", "0")

	req, err := http.NewRequest("POST", "https://api8.ocr.space/parse/image", strings.NewReader(form.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("apikey", os.Getenv("OCR_API_KEY"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var response ocrResponse
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return
	}
	// v2 splitted by \n , v3 splitted by \r\n
	words := strings.Split(response.ParsedResults[0].ParsedText, "\n")
	l := len(words)
	team.Leader = words[1]
	team.Name = words[l-2]
	team.Members = words[2 : l-3]

	return
}
