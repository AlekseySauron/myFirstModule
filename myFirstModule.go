package myFirstModule

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func Print2(path string) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("2 вариант")
	data, err := ioutil.ReadAll(resp.Body)

	type StructResult struct {
		userId    int
		id        int
		title     string
		completed bool
	}

	var result2 StructResult
	json.Unmarshal(data, &result2)

	//fmt.Println("Struct is:", result2)

	fmt.Printf("userId = %v \n", result2.userId)
	fmt.Printf("Id = %v \n", result2.id)
	fmt.Printf("title = %v \n", result2.title)
	fmt.Printf("completed = %v \n", result2.completed)

}
