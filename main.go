package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Origin struct {
	Origin string
}

func main() {
	resp, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)
	orig := new(Origin)
	err = json.NewDecoder(resp.Body).Decode(&orig)
	if err != nil {
		log.Fatalln(err)
	}
	print(orig.Origin)
}
