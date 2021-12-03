package languages

import "sort"

type Language struct {
	Name  string
	Usage int
}

func (l *Language) LanguageUsage(languages *[]Language) int {
	total := 0
	languageTotal := 0
	for _, language := range *languages {
		total += language.Usage
		if language.Name == l.Name {
			languageTotal += language.Usage
		}
	}
	percentage := (languageTotal / total) * 100
	return percentage
}

func GetTopFiveLanguages(ln *[]Language) []Language {
	SortLanguages(ln)
	if len(*ln)<=5 {
		return *ln
	}
	topFive := (*ln)[:5]
	return topFive
}

func MapToLanguage(m map[string]int) *[]Language{
	var allLanguages []Language
	for k, v := range m {
		allLanguages = append(allLanguages, Language{Name: k, Usage: v})
	}
	return &allLanguages
}

func SortLanguages(reposLanguages *[]Language){
	sort.Slice(*reposLanguages, func(i, j int) bool {
		return (*reposLanguages)[i].Usage > (*reposLanguages)[j].Usage
	})
}

func CalculateLanguagesUsage(languages []Language) *[]Language {
	mapOfStruct := make(map[string]int)
	for _, language := range languages {
		if _,found := mapOfStruct[language.Name]; found{
			mapOfStruct[language.Name] += language.Usage
		}else{
			mapOfStruct[language.Name] = language.Usage
		}
	}
	sliceOfLanguages := MapToLanguage(mapOfStruct)
	return sliceOfLanguages
}
