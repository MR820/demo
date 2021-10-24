package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var profileApi string = "http://mysql.taobao.org/monthly/"

func main() {

	url := fmt.Sprintf(profileApi)
	fmt.Println(url)

	request, err := http.NewRequest(http.MethodGet,
		url, nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("User-Agent",
		"PostmanRuntime/7.21.0")
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
	//all, err := httputil.DumpResponse(resp, true)
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(resp.Body)
	//fmt.Println(string(all))

	s := Info{}
	err = json.Unmarshal(all, &s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)

}

type Data struct {
	age             int    `json:"age"`
	educationString string `json:"educationString"`
	gender          int    `json:"gender"`
	heightString    string `json:"heightString"`
	marriageString  string `json:"marriageString"`
	memberID        int    `json:"memberID"`
	salaryString    string `json:"salaryString"`
	workCityString  string `json:"workCityString"`
}

type Info struct {
	data Data `json:"data"`
}
