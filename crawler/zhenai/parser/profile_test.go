package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParserProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("tt.html")
	if err != nil {
		panic(err)
	}
	res := ParserProfile(contents)
	for _, v := range res {
		fmt.Printf("%s\n", v)
	}

}
