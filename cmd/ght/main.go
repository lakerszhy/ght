package main

import (
	"fmt"
	"net/http"

	"github.com/lakerszhy/ght"
)

func main() {
	data, err := http.Get("https://github.com/trending/go?since=daily")
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()

	repos, err := ght.Parse(data.Body)
	if err != nil {
		panic(err)
	}

	for _, r := range repos {
		fmt.Printf("%s/%s\n", r.Owner, r.Name)
	}
}
