package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	for {
		url := os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")

		resp, _ := http.Get(url)
		defer resp.Body.Close()

		byteArray, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(byteArray))
		time.Sleep(3 * time.Second)
	}
}
