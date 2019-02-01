package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for {
		url := "http://localhost:9090"

		resp, _ := http.Get(url)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteArray))
		time.Sleep(3 * time.Second)
	}
}
