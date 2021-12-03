package main

import (
	"bufio"
	"github.com/olekukonko/tablewriter"
	"github.com/stefka-sarafska/ChaosCamp-Go/github-users-statistics/languages"
	"github.com/stefka-sarafska/ChaosCamp-Go/github-users-statistics/users"
	"log"
	"os"
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
		userLanguages := user.GetAllUserLanguages()
		calculatedLanguages := languages.CalculateLanguagesUsage(userLanguages)
		topFiveLanguages := languages.GetTopFiveLanguages(calculatedLanguages)
		languagesString := createStringOfLanguages(&topFiveLanguages)
		data = append(data, []string{user.Name, strconv.Itoa(len(user.Repositories)), languagesString, strconv.Itoa(user.Followers), strconv.Itoa(user.GetUserForks())})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"User", "Number of user repos", "Languages", "Followers", "Forks"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.AppendBulk(data) // Add Bulk Data
	table.SetCenterSeparator("|")
	table.SetRowLine(true)
	table.Render()
}

func createStringOfLanguages(lns *[]languages.Language) string {
	languagesString := ""
	for _, language := range *lns {
		languagesString += language.Name + " -> " + strconv.Itoa(language.Usage) + ", "
	}
	return languagesString
}
