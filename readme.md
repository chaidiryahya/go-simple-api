# Simple API with GoLang

### Requirement
[GoLang Version](https://golang.org/dl/) >= 1.11

### How to run
 - Clone repo
 - run in terminal / CLI : `go build && ./go-simple-api`
 - Wait for dependecy download
 - It all set when you see `starting serve on :8080`
 
 ### Sample API List
 - GET http://localhost:8080/
 - GET http://localhost:8080/get
 - GET http://localhost:8080/get_with_route_param/:user_id
 - GET http://localhost:8080/get_with_query_param?user_id=100
 - POST http://localhost:8080/post_with_json_body
 
 You can use your own browser for get request and i suggest use POSTMAN for post request. 
 
 Sample curl for post request :
 
 `curl --location --request POST 'http://localhost:8080/post_with_json_body' \
--header 'Content-Type: text/plain' \
--data-raw '{
	"yourName" : "Chaidir Yahya",
	"userID" : 12345
}'`

 postman :
 
 ![Hit with postman](https://i.ibb.co/7VbB443/Screenshot-from-2021-02-01-10-13-54.png)

 Sample userID : 100 and 101
 
 Please refer to this [doc](https://golang.org/doc/) for GoLang documentation :)

 ### Others
 I have added sample unit test. Will update for another sample
 
