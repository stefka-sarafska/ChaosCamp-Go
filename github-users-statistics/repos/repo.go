package repos

import (
	"encoding/json"
	"fmt"
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
	Languages map[string]int
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
	//var result map[string]int
	json.Unmarshal([]byte(body), &r.Languages)

	//for k, v := range result {
	//	r.Languages = append(r.Languages,languages.Language{Name: k, Usage: v})
	//}
}

//
//func SetRepoLanguages(r *Repository,userName string) {
//	//a := &r.Languages
//	//b := r.Languages
//	//print(a,b)
//	currentRepoLanguagesUrl := fmt.Sprintf(languagesRepositoriesUrl,userName,r.Name)
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
//	//var ln []languages.Language
//	for _, language := range splittedLanguages {
//		languageAndCount := strings.Split(language,":")
//		languageName := reg.ReplaceAllString(languageAndCount[0],"$1")
//		languageUsage := reg.ReplaceAllString(languageAndCount[1],"$1")
//		languageUsageToInt, err := strconv.Atoi(languageUsage)
//		if err != nil{
//			languageUsageToInt = 0
//		}
//		r.Languages = append(r.Languages,languages.Language{Name: languageName, Usage: languageUsageToInt})
//	}
//	//r.Languages=&ln
//   //return ln
//
//}
