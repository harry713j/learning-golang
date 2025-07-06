package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://oh-hello.me:5000/vinod?att=humvinod&goal=up-pradhan"

func main() {
	fmt.Print("URL Opeartions\n")

	res, _ := url.Parse(myurl)
	fmt.Println(res)
	fmt.Printf("Type: %T\n", res)
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Port())
	fmt.Println(res.Path)
	fmt.Println(res.RawQuery)

	queryparams := res.Query()

	fmt.Printf("Type of query: %T\n", queryparams)
	fmt.Println(queryparams)
	fmt.Println(queryparams["att"])

	// create a url
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "zzz.in",
		Path:     "/meow",
		RawQuery: "name=hula",
	}

	fmt.Println(partsOfURL.String())

}
