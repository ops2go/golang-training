package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// error is returned but left as a blank var
	res, _ := http.Get("http://www.github.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
