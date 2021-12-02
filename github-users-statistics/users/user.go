package users

import (
	"encoding/json"
	"fmt"
	"github.com/stefka-sarafska/ChaosCamp-Go/github-users-statistics/repos"
	"io/ioutil"
	"log"
	"net/http"
)

const userURL = "https://api.github.com/users/%s"
const userRepositoriesUrl = "https://api.github.com/users/%s/repos"
const languagesRepositoriesUrl = "https://api.github.com/repos/%s/%s/languages"

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
	if err := json.Unmarshal(body, &resultInfo); err != nil { // Parse []byte to the go struct pointer
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
	for _, repo := range u.Repositories {
		repo.SetRepoLanguages(u.Name)
	}
}

//func getRepoLanguages(r string,userName string) []languages.Language{
//	currentRepoLanguagesUrl := fmt.Sprintf(languagesRepositoriesUrl,userName,r)
//	resp, err := http.Get(currentRepoLanguagesUrl)
//	if err!=nil{
//		log.Fatal(err)
//	}
//	defer resp.Body.Close()
//
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil{
//		log.Fatal(err)
//	}
//	respBody := string(body)
//	//var foundLanguages []languages.Language
//
//	reg := regexp.MustCompile("{|}|\"")
//	splittedLanguages := strings.Split(respBody,",")
//	var ln []languages.Language
//	for _, language := range splittedLanguages {
//		languageAndCount := strings.Split(language,":")
//		languageName := reg.ReplaceAllString(languageAndCount[0],"$1")
//		languageUsage := reg.ReplaceAllString(languageAndCount[1],"$1")
//		languageUsageToInt, err := strconv.Atoi(languageUsage)
//		if err != nil{
//			languageUsageToInt = 0
//		}
//		ln = append(ln,languages.Language{Name: languageName, Usage: languageUsageToInt})
//	}
//	return ln
//}
