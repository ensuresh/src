package twitterapi

import (
	"net/http"
	"strings"
	"fmt"
	"os"
	"encoding/json"
	"twittersearch/config"
	"log"
	"twittersearch/cnckafkaapi"
)

//Variable to store access token
var App_Access_Token = config.Twitter_App_Only_AccessToken{}

//function to retrieve token to access twitter
func Get_Access(b64Token string){

	//Using http request method to POST API Token Request to Twitter API
	req, err := http.NewRequest("POST","https://api.twitter.com/oauth2/token",strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		fmt.Println("error", err)
		os.Exit(-1)
	}

	//Setting up Authorization
	req.Header.Add("Authorization", "Basic "+b64Token)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Do: ", err)
		return
	}
	defer resp.Body.Close()

	//Creating variable

	if err := json.NewDecoder(resp.Body).Decode(&App_Access_Token); err != nil {
		fmt.Println(err)
	}
}

//function to search for tweets using Twitter Standard APIs. Uses the token obtained from the previous function to obtain access to API
func Twitter_search(searchword string){

	url := fmt.Sprintf("https://api.twitter.com/1.1/search/tweets.json?q=%s&lang=en&count=100",searchword)

	req1, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Unable to connect")
	}

	req1.Header.Add("Authorization", fmt.Sprintf("Bearer %s", App_Access_Token.Access_Token))

	client := &http.Client{}
	resp1, err := client.Do(req1)
	if err != nil {
		fmt.Println(err)
	}
	defer resp1.Body.Close()
	fmt.Println(resp1.StatusCode, resp1.Status)

	var record config.Tweet_result

	if err := json.NewDecoder(resp1.Body).Decode(&record); err != nil {
		log.Println(err)
	}


	for i :=0 ; i < len(record.Statuses); i++ {
		//fmt.Println(i)
		//fmt.Printf("%+v\n" , record.Statuses[i])
		var this_tweet config.Single_Tweet_result
		this_tweet = record.Statuses[i]
		cnckafkaapi.PostMessageToKafka(this_tweet,"twitterfeed")

	}
	fmt.Println(record.SearchMetadata.Count)

}
