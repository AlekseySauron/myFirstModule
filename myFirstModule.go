package myFirstModule

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

func Print1(path string) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("1 вариант")

	var result1 map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result1)

	fmt.Printf("userId = %v \n", result1["userId"])
	fmt.Printf("Id = %v \n", result1["id"])
	fmt.Printf("title = %v \n", result1["title"])
	fmt.Printf("completed = %v \n", result1["completed"])
}
