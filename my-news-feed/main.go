package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	baseURL, _ := url.Parse("https://api.thenewsapi.com")
	baseURL.Path += "v1/news/all"

	params := url.Values{}
	params.Add("api_token", "y3KXpbHzDMy6GF1NcbKaz0YRrx4wkkK4BQvpjQtm")
	params.Add("categories", "business,tech")
	params.Add("search", "apple")
	params.Add("limit", "3")

	baseURL.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", baseURL.String(), nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
