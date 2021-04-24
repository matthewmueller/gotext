package gotext_test

import (
	"testing"

	"github.com/matthewmueller/gotext"
	"github.com/tj/assert"
)

type fntest struct {
	fn  func(...string) string
	in  string
	out string
}

func TestSingular(t *testing.T) {
	tests := []fntest{
		{gotext.Singular, "", ""},
		{gotext.Singular, "hi world", "hi world"},
		{gotext.Singular, "hi worlds", "hi world"},
		{gotext.Singular, "Hi Worlds", "Hi World"},
		{gotext.Singular, "Hi$Worlds$$$", "Hi$World$$$"},
		{gotext.Singular, "hi_worlds", "hi_world"},
		{gotext.Singular, "his_worlds", "his_world"},
		{gotext.Singular, "his", "hi"},
		{gotext.Singular, "Blogs", "Blog"},
		{gotext.Singular, "Blog", "Blog"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestPlural(t *testing.T) {
	tests := []fntest{
		{gotext.Plural, "", ""},
		{gotext.Plural, "hi worlds", "hi worlds"},
		{gotext.Plural, "hi world", "hi worlds"},
		{gotext.Plural, "Hi World", "Hi Worlds"},
		{gotext.Plural, "Hi$World$$$", "Hi$Worlds$$$"},
		{gotext.Plural, "hi_world", "hi_worlds"},
		{gotext.Plural, "his_world", "his_worlds"},
		{gotext.Plural, "his", "his"},
		{gotext.Plural, "hi", "his"},
		{gotext.Plural, "Blogs", "Blogs"},
		{gotext.Plural, "Blog", "Blogs"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestCamel(t *testing.T) {
	tests := []fntest{
		{gotext.Camel, "", ""},
		{gotext.Camel, "hi world", "hiWorld"},
		{gotext.Camel, "user id", "userID"},
		{gotext.Camel, "id user", "idUser"},
		{gotext.Camel, "id dns http", "idDNSHTTP"},
		{gotext.Camel, "http user", "httpUser"},
		{gotext.Camel, "string", "_string"},
		{gotext.Camel, "newIn", "newIn"},
		{gotext.Camel, "new", "_new"},
		{gotext.Camel, "Blog", "blog"},
		{gotext.Camel, "ASKED", "asked"},
		{gotext.Camel, "ASKED_RECIEVED", "askedRecieved"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestPascal(t *testing.T) {
	tests := []fntest{
		{gotext.Pascal, "", ""},
		{gotext.Pascal, "hi world", "HiWorld"},
		{gotext.Pascal, "user id", "UserID"},
		{gotext.Pascal, "id user", "IDUser"},
		{gotext.Pascal, "id dns http", "IDDNSHTTP"},
		{gotext.Pascal, "http user", "HTTPUser"},
		{gotext.Pascal, "string", "String"},
		{gotext.Pascal, "Blog", "Blog"},
		{gotext.Pascal, "ASKED", "Asked"},
		{gotext.Pascal, "ASKED_RECIEVED", "AskedRecieved"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}

func TestSlim(t *testing.T) {
	tests := []fntest{
		{gotext.Slim, "", ""},
		{gotext.Slim, "hi world", "hiworld"},
		{gotext.Slim, "user id", "userid"},
		{gotext.Slim, "user_id", "userid"},
		{gotext.Slim, "userid", "userid"},
		{gotext.Slim, "id user", "iduser"},
		{gotext.Slim, "id dns http", "iddnshttp"},
		{gotext.Slim, "http user", "httpuser"},
		{gotext.Slim, "string", "_string"},
		{gotext.Slim, "Blog", "Blog"},
	}
	for _, test := range tests {
		assert.Equal(t, test.out, test.fn(test.in), "%s != %s", test.in, test.out)
	}
}
