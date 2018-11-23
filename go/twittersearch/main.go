package main

import (
	"twittersearch/config"
	"twittersearch/twitterapi"
	"os"
)

func main(){
	b64Token := config.Get_Keys()
	twitterapi.Get_Access(b64Token)
	twitterapi.Twitter_search(os.Args[1])


}
