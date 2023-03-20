package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// GraphQL query
	query := `
		query {
			Workouts {
				_id
				title
			}
		}
	`

	// JWT token
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.Apuw3rKYppYA9PKrqQemkyBKufwoSJP_5GKamkSw8Mg"

	// GraphQL endpoint
	url := "http://localhost:5000/graphql"

	// GraphQL request
	requestBody, err := json.Marshal(map[string]string{
		"query": query,
	})
	if err != nil {
		fmt.Println("Error marshalling request body:", err)
		return
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error creating GraphQL request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Send GraphQL request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending GraphQL request:", err)
		return
	}
	defer resp.Body.Close()

	// Parse GraphQL response
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error parsing GraphQL response:", err)
		return
	}

	// Print response data
	fmt.Println(data)
}
