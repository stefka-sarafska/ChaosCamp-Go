package languages

type Language struct {
	Name  string
	Usage int
}

func (l *Language) languageUsage(languages *[]Language) int {
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
