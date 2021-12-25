package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func main() {
	// Parses flag to enter port number when stating application. Default port is 8080
	var port = flag.Int("port", 8080, "Port that server listens on")
	flag.Parse()

	// Assigns functions to different request paths
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", rootHandler)

	// Serves application and checks if an error occured. Listens on given port
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Println("Error serving web app.", err)
	}
}

// Handles http requests to root path /
// Sends short responce message and outputs same message to server output
func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Works")
	w.Write([]byte("Works"))
}

// Handles http requests to path /ws
// Expects upgrade to WebSocket
// Validates connection and begins message communication
// Returns if connection error occurs, header Content-Type is not "text/plain"
//  or message received is not correct type
// Maintains connection if all data is correct
func wsHandler(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket connection error.", err)
		return
	}
	defer conn.Close()

	if r.Header == nil || r.Header["Content-Type"] == nil || !contains(r.Header["Content-Type"], "text/plain") {
		log.Println("Content-Type is not correct. Given:", r.Header["Content-Type"])
		return
	}
	log.Println("Client created a WebSocket.")

	waitForRequests(conn)
}

// Waits for messages sent through initialized connection with the client
// Returns if messages are null or type is not TextMessage
// If message is not empty, replaces question marks with exclamation marks
// 	and sends result back to client
func waitForRequests(conn *websocket.Conn) {
	for {
		mType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Failed to read message.", err)
			return
		}
		if mType != websocket.TextMessage {
			log.Printf("Message was not text. Type: %d.\n", mType)
			return
		}

		message := string(data)
		if message != "" {
			log.Println("Message received:", message)
			replacedMessage := replaceQuestionMarks(message)
			conn.WriteMessage(websocket.TextMessage, []byte(replacedMessage))
		}
	}
}

// Returns new string with question marks replaced [?] with exclamation marks [!]
func replaceQuestionMarks(msg string) string {
	return strings.Replace(msg, "?", "!", -1)
}

// Checks if array of strings [arr] contains given string [str]
// Returns true if contains. Returns false if [arr] is empty or not contains [str]
func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
