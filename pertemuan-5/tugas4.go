// hapus API
package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Delete Post Response:", resp.Status)
}
