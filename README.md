# golang-repost-bot
This repository stores the code for a Twitter bot written in Golang.

This is my very first Go project. I intend to use it to learn about HTTP server and how to work with JSON in Go.

What this bot does is the following:

* Makes a GET request for a subreddit

* Gets 15 posts

* Checks which posts have 1000 upvotes or more

It has 3 functions:

getPosts(), twitterWebhook() and greaterThan()

## Functions

getPosts makes a GET request to the subreddit, gets the first 15 posts and then it tests if the upvotes are greater than 1000 and less than 100.000. After that, it uses a random function to randomly pick a post.

greaterThan is a function to check if the upvotes are greater than 1000 and less than 100.000. I did that because I consider posts with 1000 upvotes and above as "highlights", and the 100.000 value is due to a fixed post that I won't consider.

twitterWebhook if responsible for sending the request to the twitter API endpoint. It also contains the API keys and everything else needed to make the request.

## How to run it?

You must install Go, I developed this bot using the 1.14 version.

After installing go, just go to the cmd directory and run the bot with go run bot.go.
