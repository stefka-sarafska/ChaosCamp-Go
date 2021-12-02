package main

import (
	"bufio"
	"github.com/olekukonko/tablewriter"
	"github.com/stefka-sarafska/ChaosCamp-Go/github-users-statistics/languages"
	"github.com/stefka-sarafska/ChaosCamp-Go/github-users-statistics/users"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	files := os.Args[1:]
	for _, filename := range files {
		err := processFile(filename)
		if err != nil {
			log.Printf("Error processing %s : %v\n", filename, err)
		}
	}

}

func processFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var names []string
	for scanner.Scan() {
		line := scanner.Text()
		names = append(names, line)
	}
	createTable(names)
	if scanner.Err() != nil {
		return scanner.Err()
	}
	return nil
}

func createTable(userNames []string) {
	var data [][]string
	for _, name := range userNames {
		user := &users.User{Name: name}
		user.SetUserInfo()
		data = append(data, []string{user.Name, strconv.Itoa(len(user.Repositories)), "Java -> 100\nCss ->50", strconv.Itoa(user.Followers), strconv.Itoa(user.GetUserForks()), "activity"})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"User", "Number of user repos", "Languages", "Followers", "Forks", "Activity by year"})
	table.SetBorder(true)  // Set Border to false
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}

//func getAllUserLanguages(user *users.User) []languages.Language{
//	userRepos := user.Repositories
//	var languages []languages.Language
//	for _, repo := range userRepos {
//		for _, language := range repo.Languages {
//			languages = append(languages, language)
//		}
//	}
//	return languages
//}

func getTopFiveLanguages(languages []languages.Language) []languages.Language {
	sort.Slice(languages, func(i, j int) bool {
		return languages[i].Usage < languages[j].Usage
	})
	if len(languages) > 5 {
		topLanguages := languages[:len(languages)-5]
		return topLanguages
	}
	return languages
}
