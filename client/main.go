package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	r, err := http.Get("http://127.0.0.1:8080/stream")
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	reader := bufio.NewReader(r.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		fmt.Println(line)
	}
	fmt.Println("end of stream")
}
