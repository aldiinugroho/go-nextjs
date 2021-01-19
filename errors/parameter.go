package errors

import (
	"regexp"
	"strings"
)

func Parameter_tag(x string) string {
	if strings.Count(x, "/") == 2 {
		fix := lastString(strings.Split(x, "/"))
		match, _ := regexp.MatchString("[a-zA-Z ]+", fix)
		if match == false {
			return fix
		}
		return "0"
	}
	return "0"
}

func Parameter_detail(x string) string {
	if strings.Count(x, "/") == 2 {
		fix := lastString(strings.Split(x, "/"))
		match, _ := regexp.MatchString("[a-zA-Z ]+", fix)
		if match == false {
			return fix
		}
		return "0"
	}
	return "0"
}

func lastString(ss []string) string {
	return ss[len(ss)-1]
}
