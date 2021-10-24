package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	Fetch("http://album.zhenai.com/u/101285875")
}

func Fetch(url string) {

	request, err := http.NewRequest(http.MethodGet,
		url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36")
	client := http.Client{
		CheckRedirect: func(
			req *http.Request,
			via []*http.Request) error {
			//fmt.Println("Redirect:", req)
			return nil
		},
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)

}
