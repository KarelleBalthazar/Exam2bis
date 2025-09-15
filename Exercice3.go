package main

func SplitWhitesSpace(s string) []string {
	var livre []string
	var mot string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
			mot += string(s[i])
		} else {
			if mot != "" {
				livre = append(livre, mot)
				mot = ""
			}
		}
	}
	if mot != "" {
		livre = append(livre, mot)
	}
	return livre
}
