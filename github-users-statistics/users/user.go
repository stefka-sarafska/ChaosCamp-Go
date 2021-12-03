package users

import (
	"encoding/json"
	"fmt"
	"github.com/stefka-sarafska/ChaosCamp-Go/github-users-statistics/languages"
	"github.com/stefka-sarafska/ChaosCamp-Go/github-users-statistics/repos"
	"io/ioutil"
	"log"
	"net/http"
)

const userURL = "https://api.github.com/users/%s"
const userRepositoriesUrl = "https://api.github.com/users/%s/repos"

type User struct {
	Name         string
	Followers    int
	Repositories []repos.Repository
}

func (u *User) GetUserForks() int {
	forks := 0
	for _, repository := range u.Repositories {
		forks += repository.Forks
	}
	return forks
}

func (u *User) SetUserInfo() {
	currUserUrl := fmt.Sprintf(userURL, u.Name)
	resp, err := http.Get(currUserUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var resultInfo User
	if err := json.Unmarshal(body, &resultInfo); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	u.Followers = resultInfo.Followers
	u.SetUserRepositories()

}

func (u *User) SetUserRepositories() {
	currentUserRepositoryUrl := fmt.Sprintf(userRepositoriesUrl, u.Name)
	resp, err := http.Get(currentUserRepositoryUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &u.Repositories); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	for i := range u.Repositories {
		repo := &u.Repositories[i]
		repo.SetRepoLanguages(u.Name)
	}
}

func (u *User) GetAllUserLanguages() []languages.Language {
	userRepos := u.Repositories
	var userLanguages []languages.Language
	for _, repo := range userRepos {
		for i := range repo.Languages {
			language := &repo.Languages[i]
			userLanguages = append(userLanguages, *language)
		}
	}
	return userLanguages
}
