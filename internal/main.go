package main

import (
	"fmt"
	"html"
	"log"
	"net/http"


    "github.com/go-openapi/loads"
    "github.com/go-openapi/runtime/middleware"
    "github.com/vinayaknayar/go-rest-api/pkg/swagger/server/restapi"
    "github.com/vinayaknayar/go-rest-api/pkg/swagger/server/restapi/operations"
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// log.Println("Listening on localhost:8080")

	// log.Fatal(http.ListenAndServe(":8080", nil))

	//intialize swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON,"")
	if err != nil{
		log.Fatalln(err)
	}

	api := operations.NewHelloAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func ()  {
		if err := server.Shutdown(); err != nil{
			//error handle
			log.Fatalln(err)
		}	
	}()
	
	server.Port = 8080

	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)
	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)
	api.GetGopherNameHandler = operations.GetGopherNameHandlerFunc(GetGopherByName)

	//Start server which listening
	if err := server.Serve(); err != nil{
		log.Fatalln(err)
	}
	
}