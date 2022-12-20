package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var parseURL, _ = url.Parse("/parse/image")

func GetWord(imageURL string) (word string, err error) {
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

	req, err := http.NewRequest("POST", ocrClient.URL.ResolveReference(parseURL).String(), strings.NewReader(form.Encode()))
	req.Header.Set("apikey", os.Getenv("OCR_API_KEY"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := ocrClient.HTTP.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	var response ocrResponse
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return
	}
	// v2 splitted by \n , v3 splitted by \r\n
	words := strings.Split(response.ParsedResults[0].ParsedText, "\n")

	return words[len(words)-1], nil
}
