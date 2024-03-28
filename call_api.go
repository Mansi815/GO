package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	//URL of the API endpoint
	apiUrl := "https://api.example.com/data"

	//send Get request
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("Error making Get request : %s\n", err)
		return
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body : %s \n", err)
		return
	}
	fmt.Println(string(body))
}
