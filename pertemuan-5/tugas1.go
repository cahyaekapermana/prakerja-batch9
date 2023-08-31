// Tampilkan all data
package main

import (
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
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Posts:")
	for _, post := range posts {
		fmt.Printf("ID: %d\nUserID: %d\nTitle: %s\nBody: %s\n\n", post.ID, post.UserID, post.Title, post.Body)
	}
}
