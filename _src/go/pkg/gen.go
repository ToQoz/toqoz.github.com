package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Repo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Url         string `json:"html_url"`
}

func main() {
	data, err := readData()

	if err != nil {
		panic(err)
	}

	html, err := convertJsonToHtml(data)

	if err != nil {
		panic(err)
	}

	fmt.Println(html)
}

func readData() ([]byte, error) {
	wd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	_, err = os.Stat(filepath.Join(wd, "data.json"))

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(filepath.Join(wd, "data.json"))

	if err != nil {
		return nil, err
	}

	return data, nil
}

func convertJsonToHtml(data []byte) (string, error) {
	repos := []*Repo{}
	err := json.Unmarshal(data, &repos)

	if err != nil {
		return "", err
	}

	html := "<dl>\n"

	for _, repo := range repos {
		html += "  <dt><a href=\"" + repo.Url + "\">" + repo.Name + "</a></dt>\n"
		html += "  <dd>\n"
		html += "    <p>\n"
		html += "      " + repo.Description + "\n"
		html += "    </p>\n"
		html += "  </dd>\n"
	}

	html += "</dl>"

	return html, nil
}
