package Config

import (
	"log"
	"regexp"
)

func CheckLogServer(serverURL *string) bool {
	var regex *regexp.Regexp
	var err error

	if regex, err = regexp.Compile(`(\\b(https?)://)?[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]| ^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
		log.Fatal(err)
	}
	return regex.MatchString(*serverURL)
}
