# TODOs:
  - ~Go Tests~ (Re-write with is as needed)
  - ~Base64 Short URL~
  - ~Add base62 auto-generate to CREATE endpoint.~
  - javascript front end
  - Config file
  - Docker compose for Go/Redis/JS

# Testing
```
go test ./app/
```

For coverage:
```
go test ./app/ -coverprofile=coverage.out
go tool cover -html=coverage.out
```

# Resources
[Creating A Simple Web Server With Golang](https://tutorialedge.net/golang/creating-simple-web-server-with-golang/)

[Creating a RESTful API With Golang](https://tutorialedge.net/golang/creating-restful-api-with-golang/)

[How I write HTTP services after eight years](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html)

[Enable CORS in Golang](https://stackoverflow.com/a/47368811/10587086)

[How to call a REST web service API from JavaScript?](https://stackoverflow.com/questions/36975619/how-to-call-a-rest-web-service-api-from-javascript/51854096)

[Consuming API's with JavaScript for Beginners](https://dev.to/gbudjeakp/consuming-api-s-with-javascript-for-beginners-13el)

[Golang Gorilla mux with http.FileServer returning 404](https://stackoverflow.com/questions/21234639/golang-gorilla-mux-with-http-fileserver-returning-404)
