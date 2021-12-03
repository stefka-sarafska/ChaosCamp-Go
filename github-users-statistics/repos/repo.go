package repos

import (
	"encoding/json"
	"fmt"
	"github.com/stefka-sarafska/ChaosCamp-Go/github-users-statistics/languages"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const languagesRepositoriesUrl = "https://api.github.com/repos/%s/%s/languages"

type Repository struct {
	Name      string
	Forks     int
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Languages []languages.Language
}

func (r *Repository) SetRepoLanguages(userName string) {
	currentRepoLanguagesUrl := fmt.Sprintf(languagesRepositoriesUrl, userName, r.Name)
	resp, err := http.Get(currentRepoLanguagesUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var result map[string]int
	json.Unmarshal([]byte(body), &result)

for k, v := range result {
		r.Languages = append(r.Languages,languages.Language{Name: k, Usage: v})
	}
}
