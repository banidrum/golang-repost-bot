package main

import (
	"encoding/json"
	"fmt"
	"github.com/dghubble/oauth1"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type FirstJSONLevel struct {
	Data SecondJSONLevel `json:"data"`
}

type SecondJSONLevel struct {
	Children []ThirdJSONLevel `json:"children"`
}

type ThirdJSONLevel struct {
	Data FinalJSONLevel `json:"data"`
}

type FinalJSONLevel struct {
	Ups   int    `json:"ups"`
	Title string `json:"title"`
	Link  string `json:"permalink"`
}

type Post struct {
	Ups   int
	Title string
	Link  string
}

func main() {

	twitterWebhook()

}

func getPosts() string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://old.reddit.com/r/Whatcouldgowrong/.json?limit=15", nil)

	if err != nil {
		panic(err)
	}

	req.Header.Set("User-Agent", "<your-user-agent>")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	var jsonResponse FirstJSONLevel

	err = json.Unmarshal(body, &jsonResponse)

	if err != nil {
		log.Fatal(err)
	}

	var postsArray []Post

	for i := range jsonResponse.Data.Children {

		value := jsonResponse.Data.Children[i].Data.Ups

		jsonResponse.Data.Children[i].Data.Link = "https://reddit.com" + jsonResponse.Data.Children[i].Data.Link

		post := Post{Ups: jsonResponse.Data.Children[i].Data.Ups,
			Title: jsonResponse.Data.Children[i].Data.Title,
			Link:  jsonResponse.Data.Children[i].Data.Link,
		}

		if greaterThan(value) {
			postsArray = append(postsArray, post)
		}

		fmt.Println(postsArray)

	}

	jsonArray, _ := json.Marshal(postsArray)

	fmt.Println(string(jsonArray))

	rand.Seed(time.Now().Unix())

	n := rand.Int() % len(postsArray)

	selectedPost := postsArray[n]

	selectedPostJSON, _ := json.MarshalIndent(selectedPost, "", "\t")

	fmt.Println("Selected post -> ", string(selectedPostJSON))

	return string(selectedPostJSON)

}

func greaterThan(value int) bool {
	//Adding a < 100000 condition will probably fix the fixed post issue
	if value >= 1000 && value < 100000 {
		return true
	}
	return false
}

func twitterWebhook() {

	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))

	token := oauth1.NewToken(os.Getenv("API_TOKEN"), os.Getenv("API_TOKEN_SECRET"))

	client := config.Client(oauth1.NoContext, token)

	redditPost := getPosts()

	form := url.Values{}
	form.Add("status", redditPost)

	req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/statuses/update.json", strings.NewReader(form.Encode()))

	fmt.Println(req)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)

	fmt.Println("status response", string(body))
}
