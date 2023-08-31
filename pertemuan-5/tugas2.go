// tampilkan data dengan id 3
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
	url := "https://jsonplaceholder.typicode.com/posts/3"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var post Post
	err = json.NewDecoder(resp.Body).Decode(&post)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Post with ID 3:\nID: %d\nUserID: %d\nTitle: %s\nBody: %s\n",
		post.ID, post.UserID, post.Title, post.Body)
}
