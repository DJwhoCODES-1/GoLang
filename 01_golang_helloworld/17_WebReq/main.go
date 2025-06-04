package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println(`Web Requests using "net/http"`)

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error in getting response:", err)
		return
	}
	defer res.Body.Close()

	// fmt.Printf("%T\n", res)

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return
	}

	// fmt.Println(data) // data received in bytes

	fmt.Println(string(data))
}
