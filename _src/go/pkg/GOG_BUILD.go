package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"html_url"`
	Language    string `json:"language"`
}

func main() {
	data, err := readData()

	if err != nil {
		panic(err)
	}

	html, err := convertJsonToHtml(data)

	if err != nil {
		fmt.Println(string(data))
		fmt.Println(err)
		os.Exit(1)
	}

	ioutil.WriteFile("./index.html", []byte(html), 0755)
	os.Exit(0)
}

func readData() ([]byte, error) {
	resp, err := http.Get("https://api.github.com/users/ToQoz/repos?type=public&per_page=100")

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

func convertJsonToHtml(data []byte) (string, error) {
	repos := []*Repo{}
	err := json.Unmarshal(data, &repos)

	if err != nil {
		return "", err
	}

	html := "<section>\n"
	html += "<h2>Golang package I created</h2>\n"
	html += "<dl>\n"

	for _, repo := range repos {
		if repo.Language == "Go" {
			html += `<dt><a href="` + repo.Url + `">` + repo.Name + "</a></dt>\n"
			html += "<dd><p>" + repo.Description + "</p></dd>\n"
		}
	}

	html += "</dl>\n"
	html += "</section>"

	return html, nil
}
