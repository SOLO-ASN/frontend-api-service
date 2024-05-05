package retriever

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"

	"github.com/g8rswimmer/go-twitter"
)

func twitter2() {
	url := "https://api.twitter.com/1.1/users/show.json?screen_name=TwitterDev"

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Set the authorization header using the API key
	req.Header.Set("Authorization", "Bearer AAAAAAAAAAAAAAAAAAAAAExNtgEAAAAAhS%2B4%2FuMez2QA0IPx03GPfFnPtCs%3DaDAt47w0djF7FdoWRxlx59eV96PM84I4qxPhQoBZfMp9mWKs87")

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Close the response body when done
	defer resp.Body.Close()

	// Decode the response JSON into a struct
	var user struct {
		Name           string `json:"name"`
		Description    string `json:"description"`
		FollowersCount int    `json:"followers_count"`
	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		panic(err)
	}

	// Print the user information
	fmt.Printf("Name: %s\nDescription: %s\nFollowers: %d\n", user.Name, user.Description, user.FollowersCount)
}

type authorize struct {
	Token string
}

func (a authorize) Add(req *http.Request) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
}

/**
	In order to run, the user will need to provide the bearer token.
**/
func Twitter1() {
	token := flag.String("token", "", "AAAAAAAAAAAAAAAAAAAAAExNtgEAAAAAhS%2B4%2FuMez2QA0IPx03GPfFnPtCs%3DaDAt47w0djF7FdoWRxlx59eV96PM84I4qxPhQoBZfMp9mWKs87")
	flag.Parse()
	fmt.Println()
	tweet := &twitter.Tweet{
		Authorizer: authorize{
			Token: *token,
		},
		Client: http.DefaultClient,
		Host:   "https://api.twitter.com",
	}
	fieldOpts := twitter.TweetFieldOptions{
		Expansions:  []twitter.Expansion{twitter.ExpansionEntitiesMentionsUserName, twitter.ExpansionAuthorID},
		TweetFields: []twitter.TweetField{twitter.TweetFieldCreatedAt, twitter.TweetFieldConversationID, twitter.TweetFieldAttachments},
	}

	lookups, err := tweet.SampledStream(context.Background(), fieldOpts)
	fmt.Println(lookups)
	var tweetErr *twitter.TweetErrorResponse
	switch {
	case errors.As(err, &tweetErr):
		printTweetError(tweetErr)
	case err != nil:
		fmt.Println(err)
	default:
		for _, lookup := range lookups {
			printTweetLookup(lookup)
			fmt.Println(1)
		}
	}
}

func printTweetLookup(lookup twitter.TweetLookup) {
	enc, err := json.MarshalIndent(lookup, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enc))
}

func printTweetError(tweetErr *twitter.TweetErrorResponse) {
	enc, err := json.MarshalIndent(tweetErr, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(enc))
}
