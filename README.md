# bsms
Send upto 3 sms messages to a number after shortening URLs in the text using bitly

## Steps to run the code
1.  Go version ```1.11``` and upwards is required to build this project. Please check your go version using the command ```go version```
2.  I have used ```go1.11.5``` to create the project and distributed ```go.mod``` which contains all the dependencies and can run a project outside the ```$GOPATH```.
3.  Create a directory structure of the form ```burst/github.com/sovikc```. I have used github code layout.
4.  Next ```cd``` to the directory ```sovikc``` and run the command ```git clone https://github.com/sovikc/bsms.git```
5.  Now the directory structure would be ```burst/github.com/sovikc/bsms```
6.  There is a ```.env``` file inside the root ```bsms``` directory (it would be a hidden file), where the values for env variables ```BITLY_GENERIC_ACCESS_TOKEN, BITLY_GROUP_GUID, API_KEY, API_SECRET``` need to be provided.
7.  I have used the latest version (```v4```) of ```Bitly API```. It requires a ```generic access token``` and a ```group_guid``` to shorten a URL. 
8.  If the ```BITLY_GROUP_GUID``` is not supplied in the ```.env```, the code would handle that and make an extra API call to fetch the ```group_guid``` once.
9.  I have also distributed the bundled javascript files along with the index.html in the ```bsms/static``` directory. The client-side development files are inside ```bsms/static``` directory 
10. Now ```cd``` to directory ```bsms``` and run the command ```go run main.go```.
11. This would download all the dependencies and start the server on port 8000.
12. Open a browser and type ```localhost:8000``` and you should be able to see the UI.


## User Interface
![alt text](/Screenshot_of_SMS_Form.png)
