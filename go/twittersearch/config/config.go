package config

import (
	"os"
	"encoding/base64"
	"fmt"
)

//struct for storing access token from Twitter API
type Twitter_App_Only_AccessToken struct {
	Token_Type string `json:"token_type"`
	Access_Token string `json:"access_token"`
}

//struct for storing Tweet Search results. Results are stored in an array.
type  Tweet_result struct {
	Statuses []struct {
		CreatedAt string `json:"created_at"`
		ID        int64  `json:"id"`
		IDStr     string `json:"id_str"`
		Text      string `json:"text"`
		Truncated bool   `json:"truncated"`
	} `json:"statuses"`
	SearchMetadata struct {
		CompletedIn float64 `json:"completed_in"`
		MaxID       int64   `json:"max_id"`
		MaxIDStr    string  `json:"max_id_str"`
		Query       string  `json:"query"`
		RefreshURL  string  `json:"refresh_url"`
		Count       int     `json:"count"`
		SinceID     int     `json:"since_id"`
		SinceIDStr  string  `json:"since_id_str"`
	} `json:"search_metadata"`
}

//struct for just the data needed to be stored in Kafka for single tweet result
type  Single_Tweet_result struct {
		CreatedAt string `json:"created_at"`
		ID        int64  `json:"id"`
		IDStr     string `json:"id_str"`
		Text      string `json:"text"`
		Truncated bool   `json:"truncated"`
}

//This function is to get the Consumer_key and Consumer_Secret and
func Get_Keys()(string){

	//Obtaining Twitter App Only Access Consumer Key and Consumer Secret from Environment Variable
	consumer_key := os.Getenv("CONSUMER_KEY")
	consumer_secret := os.Getenv("CONSUMER_SECRET")

	//Check for Keys
	if consumer_key == "" {
		fmt.Println("Consumer Key for Twitter API not configured")
		os.Exit(-1)
	}
	if consumer_secret == "" {
		fmt.Println("Consumer Secret for Twitter API not configured")
		os.Exit(-1)
	}

	//Encoding to 64 bit to pass to Http Request
	b64Token := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", consumer_key, consumer_secret)))
	return b64Token
}

