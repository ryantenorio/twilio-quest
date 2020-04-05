# Part 1 Basic Training

- Using environment variables
- Setting up a basic web server in Golang with Gin
  - Running the server on the correct port
  - Setting up a handler for /hello
  - Beginning to build a Twilio interface

 ```
package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.Run(":5000")
}
```

- POST localhost:5000/hello with Postman
```
<Response>
    <say>Hello, World!</say>
</Response>
```

- Getting ngrok setup
```
ngrok by @inconshreveable                                                                               (Ctrl+C to quit)                                                                                                                        Session Status                online                                                                                    Session Expires               7 hours, 59 minutes                                                                       Version                       2.3.35                                                                                    Region                        United States (us)                                                                        Web Interface                 http://127.0.0.1:4040                                                                     Forwarding                    http://6711f76f.ngrok.io -> http://localhost:5000                                         Forwarding                    https://6711f76f.ngrok.io -> http://localhost:5000                                                                                                                                                                Connections                   ttl     opn     rt1     rt5     p50     p90                                                                             2       0       0.02    0.01    0.00    0.00                                                                                                                                                                      HTTP Requests                                                                                                           -------------                                                                                                                                                                                                                                   POST /hello                    200 OK
```