# :zap: QuickStart
```
git clone git@github.com:KGB33/url-shortener.git
cd url-shortener
docker-compose up
```
Then navigate to `0.0.0.0:8080` in your browser.

# Features
  - RestAPI
  - Vue.js biased front-end.
  - Unit & Integration tests
  - Optional auto-generated Base64 Shortened URL
  - Easy to spin up using Docker Compose


# :turtle: SlowStart

Requirements
  - Go
  - Git
  - Redis

```
export DATABASE_URL="127.0.0.1:6379"
export DATABASE_PASS=""
export DATABASE_ID="0"
export PORT=":8080"
git clone git@github.com:KGB33/url-shortener.git
cd url-shortener
redis-server --daemonize yes
go run .
```

Then navigate to `127.0.0.1$PORT` in your browser.

# Testing
Run the following commands from the top-level directory of this repository.
Also, ensure that redis is running.

```
go test ./...
```
Or, if you want coverage:
```
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

# API Reference

## Json Objects
### Url
```
{ "ShortUrl": type:string, "DestUrl": type:string }
```
### Error
```javascript
{"Error": type:string}
```

## CREATE
Endpoint: `/api/v1/c`
  - Method: POST
  - Body: Url

Response:
  - On Success: 201, Url
  - On Fail   : 500 or 400, Error

Example Response:
```
{"ShortUrl": "FU4yMyyYZnF", "DestUrl": "https://golang.org/"}
```

## GET
Endpoint: `/api/v1/r`
  - Method: GET

Response:
  - On Success: 200, array[Url]
  - On Fail   : 500, {"Error": string}

Example Response:
```
[
  {
    "ShortUrl": "FU4yMyyYZnF",
    "DestUrl": "https://golang.org/"
  },
  {
    "ShortUrl": "source",
    "DestUrl": "https://github.com/KGB33/url-shortener"
  }
]
```

# UPDATE
Endpoint: `/api/v1/u/<ShortUrl>`
  - Method: PUT
  - Body: Url

Response:
  - On Success: 201, Url
  - On Fail   : 500 or 400, Error

Example:
```
PUT /api/v1/u/FU4yMyyYZnF
  Body: {"ShortUrl": "New Short URL", "DestUrl": "https://golang.org/"}

// Response
{"ShortUrl": "New Short URL", "DestUrl": "https://golang.org/"}
```

## DELETE
Endpoint: `/api/v1/d/<ShortUrl>`
  - Method: DELETE

# Resources
[Creating A Simple Web Server With Golang](https://tutorialedge.net/golang/creating-simple-web-server-with-golang/)

[Creating a RESTful API With Golang](https://tutorialedge.net/golang/creating-restful-api-with-golang/)

[How I write HTTP services after eight years](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html)

[Enable CORS in Golang](https://stackoverflow.com/a/47368811/10587086)

[How to call a REST web service API from JavaScript?](https://stackoverflow.com/questions/36975619/how-to-call-a-rest-web-service-api-from-javascript/51854096)

[Consuming API's with JavaScript for Beginners](https://dev.to/gbudjeakp/consuming-api-s-with-javascript-for-beginners-13el)

[Golang Gorilla mux with http.FileServer returning 404](https://stackoverflow.com/questions/21234639/golang-gorilla-mux-with-http-fileserver-returning-404)
