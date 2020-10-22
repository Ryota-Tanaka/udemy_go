package main

import (
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http;//localhost:18888")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	//log.Println(string(body))
	log.Println("Status:", resp.Status)
}
