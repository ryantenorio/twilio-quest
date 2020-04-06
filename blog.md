# TwilioQuest: Passing Basic Training with Go!

Are you interested in The Twilio Community Hackathon but not sure where to start? Maybe TwilioQuest is the place for you! 

## What is TwilioQuest?

[PHOTO OF TWILIOQUEST]

[TwilioQuest](https://www.twilio.com/quest) is a neat role-playing game made by Twilio to teach a variety of technical skills through missions. There are many missions that are focused on interacting with the Twilio API, but the game includes modules that are applicable to any developer, such as improving your command line skills and [contributing to open source software](https://www.twilio.com/quest/learn/open-source) on GitHub. As I flush out some ideas for the hackathon, I thought it wood be a good idea to start learning the API using TwilioQuest.

[PHOTO OF MY COOL GUY]

This hackathon is awesome because of the support being provided to participants (including free API credits! score!) and a category specifically for helping solve communication challenges as a result of the COVID-19 pandemic. I have never been a  strong Javascript developer, and saw this as a unique opportunity to learn a little bit of [Javascript](https://dev.to/ryantenorio/comment/nca8) and work on my Golang, which will be the focus of this post! 

## On to the Coding!

Twilio has official SDKs for lots of languages like Python and Javascript, but not for Go yet. There are some great libraries out there already on GitHub, but for the purpose of really learning the API I thought it was important to roll my own simple library to make it through the TwilioQuest missions. In this post, we are going to go over some of the code I wrote to get through the Basic Training mission in TwilioQuest.

## Setting up a basic webserver in Go with Gin

Go has a great standard library with the web in mind since day one. We could use the raw `http` library here but I have opted to use [Gin](https://github.com/gin-gonic/gin). There are many reasons to use or not use web frameworks, but ultimately when I sit down to write side projects I take into consideration what my goals are. In this case, my goals are to practice my writing, learn some of the Twilio API, and see if we can make some tools that can make a difference. Using Gin here to speed-up some of my web server setup is a tool towards that end! If you are following along but want to use just the standard library, the offical documentation has a pretty [good web server tutorial](https://golang.org/doc/articles/wiki/) to get you started.

One of the first tasks in Basic Training is to write a web server that returns `hello world` when you visit it. To accomplish this with Gin, all we need to do is create a new gin engine, tell the engine what to return when someone visits us, and run it! 

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

The default engine, which we are calling `router` here, sets up some basic logging middleware and recovery middleware to help handle server (500 style) errors and give you some sensible logging defaults to help with debugging in the future. We then define what happens when someone makes a GET call to our web server. We are choosing to call our engine `router` because at the basic level, most of the work we want to do with this variable is defining what to do when a user visits the different routes available on our web server. The first argument is the path and the second argument is a function of type `gin.HandlerFunc`. In the future we will write our functions elsewhere to organize the code better. This function writes the string `hello world` to send back as the response, along with a 200 HTTP Status Code (OK!). Lastly, we tell the engine to run and pass the argument `:5000` which means to run as `localhost` on port 5000.

## Accessing environment variables
The basic training mission includes some tasks that involve providing your Twilio API credentials. The game comes with a build-in IDE that supports Javascript, and these tasks are to allow you to easily access those credentials througout the game. Since we are writing in Go outside of the game, we want to make sure we are still able to access those variables. One safe way to do so is to never include your credentials in any of your files, and instead access them through environment variables. One way to do this in Go is through the `os` standard library:

```
// initialize the twilio credentials. We will need this later!
TwilioAccountSID := os.Getenv("TWILIO_ACCOUNT_SID")
TwilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")
```
You don't need to use these credentials to pass basic training, but we will want to use something like this so that we can interact with the live API later!

## Writing TwiML using Go structs and tags
TwiML is the markdown language that is used to communicate actions to the Twilio API for features like programmable voice and programmable sms. If you are writing applications with a supported SDK, the TwiML is mostly written for you. It is very xml-like and we can use that to our advantage in Go. Go structs allow you to add tags to explicitly define how a field should be represented in other formats, such as json or XML. 

To pass basic training, we have to create an endpoint that returns TwiML with say tags. To that end, we will make a struct that will represent a TwiML Response, setup a handler to send the Response, and then hook up the handler to response to the `/hello` route defined in the mission.

```
// Response is a representation of TwiML instructions
type Response struct {
	Says []Say `xml:"say"`
}
```

Our struct is called Response to match the TwiML terminology. We use the tags feature to let Go know that if we need to represent this struct in an XML format (which will make the struct TwiML compatible as well) how to name each field.

## Sending TwiML!

```
func hello(c *gin.Context) {
	// return back a say tag
	response := Response{
		Says: []Say{"Hello, World!", "Welcome to my Twilio App!"},
	}
	c.XML(200, response)
}
```

To define a handler, your function needs to take in a `*gin.Context` as its sole parameter. This context is "the most important part of Gin" according to the source code and is used to pass variables between different layers of middleware and houses many of the tools you need to render content back to the user. In this handler `hello`, we are creating a response object and telling our context to serialize the object as XML (using the tags we made earlier as hints to the compiler on how we want it!).

Now we just need to setup a route!

```
router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
router.POST("/hello", hello)
```

The first line is how we setup a route earlier with an inline function. The second line defines the function to link to the route `/hello` to be the hello function we wrote. The line begins with `router.POST` to tell our webserver to only route responses to this handler if the user uses a POST request, which is one of the requirements for one of our basic training mission tasks. 

Once that's all hooked up, you can run your server and test the path!

```
go build
./twilio-quest
```

And then in another terminal or using postman:
```
curl --location --request POST 'http://localhost:5000/hello'

<Response>
    <say>Hello, World!</say>
    <say>Welcome to my Twilio App!</say>
</Response>
```

We now have a running web server, a route to receive requests, and learned a way to start working with TwiML within Go! 

## What next?
The basic training mission includes some other good tasks to go through. One tasks is to install ngrok for your environment so you can assign a temporary public URL to route traffic to your test web server running on `localhost`. This is helpful if you want to do some quick testing that requires an internet endpoint or want to use a website like [postwoman.io](http://postwoman.io/) to test your routes.

As we go through TwilioQuest (and start thinking about what we want to do for the April 2020 Hackathon!), we will likely find good opportunities to start refactoring our code. Some likely tasks include moving all of our Twilio-related code into a single package that we can re-use and share later, support for configuration files, and tests. 

In the next article, we will take a look at the SMS API specifically and think through some ways we can creatively use it! Follow me here on dev.to or on [twitter](https://twitter.com/leeto) to get updated when the post drops! If there are other topics about Go, such as some tips on getting started with Go Modules (AKA Go's package.json) or APIs you want to explore, let me know below!