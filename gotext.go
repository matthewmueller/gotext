// Package gotext is a wrapper over matthewmueller/text for generating
// go-compatible code. This means replacing keywords with valid alternatives
// and handling common acronyms that golint warns about.
//
// One of the goals for this package is to have orthogonal functions. For
// example, text.Snake("Hi World") should return `Hi_World`, not `hi_word`. This
// can be challenging to pull off at times and I don't think we're quite there
// yet. Once we feel like we've reached the correct balance, I think it makes
// sense to pull out this package and make it available for everyone.
package gotext

import (
	"strings"

	"github.com/matthewmueller/text"
)

// Unreserved returns an unreserved version of the word
func Unreserved(s string) string {
	if builtins[s] != "" {
		return builtins[s]
	}
	return s
}

// Space case
func Space(s ...string) string {
	return Unreserved(text.Space(strings.Join(s, " ")))
}

// Lower case
func Lower(s ...string) string {
	return Unreserved(text.Lower(strings.Join(s, " ")))
}

// Upper case
func Upper(s ...string) string {
	return Unreserved(text.Upper(strings.Join(s, " ")))
}

// Camel case
func Camel(s ...string) string {
	a := strings.Split(text.Space(strings.Join(s, " ")), " ")
	for i, w := range a {
		if i == 0 {
			a[i] = strings.ToLower(a[i])
			continue
		}
		// avoid initialisms
		uw := text.Upper(w)
		if initialisms[uw] {
			a[i] = uw
		} else {
			a[i] = strings.Title(strings.ToLower(w))
		}
	}
	return Unreserved(strings.Join(a, ""))
}

// Pascal case
func Pascal(s ...string) string {
	a := strings.Split(text.Space(strings.Join(s, " ")), " ")
	for i, w := range a {
		// avoid initialisms
		uw := text.Upper(w)
		if initialisms[uw] {
			a[i] = uw
		} else {
			a[i] = strings.Title(strings.ToLower(w))
		}
	}
	return Unreserved(strings.Join(a, ""))
}

// Snake case
func Snake(s ...string) string {
	return Unreserved(text.Snake(strings.Join(s, " ")))
}

// Slug case
func Slug(s ...string) string {
	return Unreserved(text.Slug(strings.Join(s, " ")))
}

// Title case
func Title(s ...string) string {
	return Unreserved(text.Title(strings.Join(s, " ")))
}

// Short case
func Short(s ...string) string {
	return Unreserved(text.Short(strings.Join(s, " ")))
}

// Singular string
func Singular(s ...string) string {
	return Unreserved(text.Singular(strings.Join(s, " ")))
}

// Plural string
func Plural(s ...string) string {
	return Unreserved(text.Plural(strings.Join(s, " ")))
}

// Slim string
func Slim(s ...string) string {
	return Unreserved(strings.Join(strings.Split(text.Space(strings.Join(s, " ")), " "), ""))
}
