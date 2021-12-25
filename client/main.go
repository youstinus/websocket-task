package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

// Client main method
// Requests WebSocket with given address and port number and path
// Waits for console input
// Reads one line and sends to websocket
// Waits for responce message back from server and prints out
func main() {
	address := "ws://localhost:8080/ws"

	// Add header with Content-Type
	header := http.Header{}
	header.Set("Content-Type", "text/plain")

	// Tries to open a WebSocket
	c, _, err := websocket.DefaultDialer.Dial(address, header)
	if err != nil {
		log.Println("Could not create a WebSocket.", err)
		return
	}
	defer c.Close()

	// Scans console input as text and sends to a WebSocket in a different goroutine
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		for {
			scanner.Scan()
			text := scanner.Text()
			bytes := []byte(text)
			c.WriteMessage(websocket.TextMessage, bytes)
		}
	}()

	// Waits for the response from server and prints out to console
	for {
		_, bytes, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading message.", err)
			return
		}
		fmt.Println(string(bytes))
	}
}
