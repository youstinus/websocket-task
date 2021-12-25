## Websocket Test API


### How to execute

First run server using: `make run-server`.

Then run many client instances with: `make run-client`.

Then start typing in clients, and when you enter, text will be sent to server and returned as response.

Try yourself. Replaces question marks `?` with other symbol `!`.

All methods are in `main.go` file

App can be compiled with `go build main.go`

App can be tested with `go test`

App can be executed with `go run main.go`

Optional parameter can be added port. Example `go run main.go -port=8080` or `server -port 8080`

Type `server --help` for help


#### How it works

The server opens given or default port and listens for HTTP requests.

Created two endpoints. One is root to check if the server works.

Second is `/ws`. Expects for connection upgrade from HTTP to a WebSocket and waits for messages.

Received messages validated if it is text and all question marks replaced with exclamation marks.

Processed message is sent back to the client.

The connection closes when received not a valid message of client closes the connection.

The server should work in any environment



Client initializes http protocol change into WebSocket protocol.

The client sends all messages entered from console.

Receives processed messages and prints out to console.



I have chosen these approaches because they are documented by golang developers.

gorilla/websocket is mostly used in WebAPIs.

websocket.DefaultDialer.Dial method inside client automatically initializes http protocol change to WebSocket.

Unit tests not well done. It is possible to cover more situations.


#### References used

https://yalantis.com/blog/how-to-build-websockets-in-go/

https://quii.gitbook.io/learn-go-with-tests/build-an-application/websockets

https://stackoverflow.com/questions/47637308/create-unit-test-for-ws-in-golang

https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04

https://stackoverflow.com/questions/52964304/websocket-set-protocol-and-origin

https://www.geeksforgeeks.org/http-headers-content-type/

https://codeburst.io/unit-testing-for-rest-apis-in-go-86c70dada52d

https://www.youtube.com/watch?v=ttKgBttwzrg

https://stackoverflow.com/questions/55536439/how-can-i-upgrade-a-client-http-connection-to-websockets-in-golang-after-sending

https://golang.org/doc/articles/wiki/

https://stackoverflow.com/questions/55094133/how-to-send-some-event-update-from-http-handler-to-a-websocket-handler

