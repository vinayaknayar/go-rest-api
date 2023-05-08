# Go REST API Example

This is a simple REST API built with Go, following the tutorial at [https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm](https://dev.to/aurelievache/learning-go-by-examples-part-2-create-an-http-rest-api-server-in-go-1cdm).

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:

> git clone https://github.com/vinayaknayar/go-rest-api.git

2. Change into the directory:


3. Run the application:

> go run internal/main.go


4. Test the API by sending requests to the following endpoints:

 - `GET /healthz`: Returns a 200 OK status if the server is running.
 - `GET /hello/{name}`: Returns a personalized greeting for the specified name.
 - `GET /gopher/{name}`: Returns an image of a gopher with the specified name.

