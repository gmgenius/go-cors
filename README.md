# gocors
Go middleware for allowing Cross Origin Resource Sharing (CORS). Support [pre-flight request](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS#Preflighted_requests)

## installation
To install go-cors you can use the dep \
`dep ensure -add github.com/gmgenius/gocors`

## Example
In the example below, we are creating a server and use cors.HandleCors struct as our root path HTTP handler. The HandleCors struct implements the ServeHTTP method.
```go
 package main
 
 import (
 	"fmt"
 	"github.com/gmgenius/gocors"
 	"log"
 	"net/http"
 	"time"
 )
 
 func main() {
 	//initialize server mux to use
 	serverMux := http.NewServeMux()
 
 	//handle the root path. we use our gocors HandleCors
 	serverMux.Handle("/", gocors.HandleCors{
 		AllowOrigin:  "*",
 		AllowHeaders: "custom-heading, content-type",
 		HttpHandler: http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
 			//set headers for our response writer
 			writer.Header().Add("Content-Type", "application/json")
 			fmt.Println(request.Method)
 			fmt.Fprintln(writer, `{"name":"George"}`)
 		}),
 	})
 
 	//initialize our server
 	server := &http.Server{
 		Handler:      serverMux,
 		WriteTimeout: 15 * time.Second,
 		ReadTimeout:  15 * time.Second,
 		Addr:         ":3000",
 	}
 
 	fmt.Println("Server up and running")
 	//start the server
 	log.Fatal(server.ListenAndServe())
 }
```

## Initializing the HandleCors Struct
The struct has four properties, but only one is required to be provided during initialization.
* AllowOrigin (not-required); default value = "*"
* AllowMethods (not-required); default value = "GET, POST"
* AllowHeaders (not-required); default value = "Content-Type"
* HttpHandler (required)

```go
type HandleCors struct {
	AllowOrigin  string       //origin to be allowed. Domains
	AllowMethods string       //methods to be allowed in cors
	AllowHeaders string       //headers to be allowed in cors
	HttpHandler  http.Handler //handler to be called to continue with request processing
}
```

## Copyright
See LICENCE file for more details.

