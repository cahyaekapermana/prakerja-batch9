// Simpan data api
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	newPost := Post{
		UserID: 1,
		ID:     101,
		Title:  "New Post",
		Body:   "This is a new post.",
	}

	payload, err := json.Marshal(newPost)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Create Post Response:", resp.Status)
}
