package regex

import (
	"regexp"
)

//Subexp returns the subexp with name subexp from target or "" if it does not exist.
func Subexp(r *regexp.Regexp, target string, subexp string) (val string) {
	matches := r.FindStringSubmatch(target)
	for i, name := range r.SubexpNames() {
		if i > len(matches) {
			return
		}
		if name == subexp {
			return matches[i]
		}
	}
	return
}

//SubexpMap returns a map with keys of form subexpName -> value
func SubexpMap(r *regexp.Regexp, target string) map[string]string {
	matches := r.FindStringSubmatch(target)
	m := make(map[string]string, len(r.SubexpNames()))
	for i, name := range r.SubexpNames() {
		if i > len(matches) { //no more matches
			return m
		}
		if i == 0 { //first subexp is the target
			continue
		}
		m[name] = matches[i]
	}
	return m
}
