package main

import (
	"context"
	"fmt"
	"log"

	"github.com/levigross/grequests/v2"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	var users []User
	url := "https://jsonplaceholder.typicode.com/users"
	resp, err := grequests.Get(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}
	// var out map[string]any
	if err := resp.JSON(&users); err != nil {
		log.Fatal(err)
	}
	fmt.Println(users)
}
