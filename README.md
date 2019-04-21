# bsms
Send upto 3 sms messages to a number after shortening URLs in the text using bitly

## Steps to run the code
* I have used Go modules for dependencies, which is available from go version 1.11 onwards. I have used go 1.11.5 to create the project and distributed go.mod which contains all the dependencies.
* I have also distributed the client-side javascript files along with the index.html in the static directory.
* The .env file needs bitly generic access token, group_guid, and burst sms API Key and Secret. 
* The command ```go run main.go``` would start the server on port 8000.

