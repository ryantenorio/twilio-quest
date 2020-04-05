# TwilioQuest: Passing Basic Training with Go!

Are you interested in The Twilio Community Hackathon but not sure where to start? Maybe TwilioQuest is the place for you! 

## What is all this?!

TwilioQuest is a neat role-playing game made by Twilio to teach a variety of technical skills through missions. There are many missions that are focused on interacting with the Twilio API, but the game includes modules that are applicable to any developer, such as improving your command line skills and [contributing to open source software](https://www.twilio.com/quest/learn/open-source) on GitHub. 

This hackathon is awesome because of the support being provided to participants (including free API credit! score!) and a category specifically for helping solve communication challenges as a result of the COVID-19 pandemic. I have never been a really strong Javascript developer, and saw this as a unique opportunity to learn a little bit of [Javascript](https://dev.to/ryantenorio/comment/nca8) and work on my Golang, which will be the focus of this post!

## On to the coding!

Twilio has official SDKs for lots of languages like Python and Javascript, but nothing for Go. There are some great libraries out there, but for the purpose of really learning the API I thought it was important to roll my own simple library to make it through the TwilioQuest missions. In this post, we are going to go over some of the code I wrote to get through the Basic Training mission in TwilioQuest.

## Setting up a basic webserver in Go with Gin

Go has a great standard library that had the web in mind with day one. We could use the raw `http` library here but I have opted to use Gin. There are many reasons to use or not use web frameworks, but ultimately when I sit down to write side projects I take into consideration what are my actual goals. In this case, my goals are to practice my writing, learn some of the Twilio API, and see if we can make some tools that can make a difference. Using Gin here to speed-up some of my web server setup is a tool towards that end. If you are following along but want to use just the standard library, the offical documentation has a pretty [good web server tutorial](https://golang.org/doc/articles/wiki/) to get you started! 

One of the first tasks is to write a web server that returns `hello world` when you visit it. To accomplish this with Gin, all we need to do is create a new gin engine, tell the engine what to return when someone visits us, and run it! 

```
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
    router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	router.Run(":5000")
}
```

The default engine, which we are calling `router` here, sets up some basic logging middleware and recovery middleware to help handle server (500 style) errors and give you some sensible logging defaults to help with your debugging in the future. We then define what happens when someone makes a GET call to our web server. We are choosing to call our engine `router` because at the basic level, most of the work we want to do with this variable is defining what to do when a user visits the different routes available on our web server. The first argument is the path and the second argument is a function of type `gin.HandlerFunc`. In the future we will write our functions elsewhere to organize the code better. This function writes the string `hello world` to send back as the response, along with a 200 HTTP Status Code (OK!). Lastly, we tell the engine to run and pass the argument `:5000` which means to run as `localhost` on port 5000.

## Accessing environment variables
The basic training missions includes some tasks that involve providing your Twilio API credentials. The game comes with a build-in IDE that supports Javascript, and these tasks are to allow you to easily access those credentials througout the game. Since we are writing in Go outside of the game, we want to make sure we are still able to access those variables. One safe way to do so is to never include your credentials in any of your files, and instead access them through environment variables. One way to do this in Go is through the `os` standard library:

```
// initialize the twilio credentials. We will need this later!
TwilioAccountSID := os.Getenv("TWILIO_ACCOUNT_SID")
TwilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
```
You don't need to use these credentials to pass basic training, but we will want to use something like this so that we can interact with the live API later!

## Writing TwiML using Go structs and tags
TwiML is the markdown language that is used to communicate actions to the Twilio API. It is very xml-like and we can use that to our advantage in Go.

## Setting up ngrok
