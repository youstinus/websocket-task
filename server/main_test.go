package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestRootHandler(t *testing.T) {
	request, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(rootHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "OK status code is expected")
	assert.Equal(t, "Works", response.Body.String(), "Incorrect string found")
}

func TestWsHandlerHttpBadRequest(t *testing.T) {
	request, err := http.NewRequest("GET", "/ws", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(wsHandler)
	handler.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code, "BadRequest status code is expected")
}

func TestWsHandlerContentTypeIncorrect(t *testing.T) {
	handler := http.HandlerFunc(wsHandler)
	s := httptest.NewServer(handler)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Create incorrect header
	header := http.Header{}
	header.Set("Content-Type", "application/json")
	// Connect to the server
	_, _, err := websocket.DefaultDialer.Dial(u, header)
	if err != nil {
		t.Fatalf("Could not open a ws connection on %s %v", u, err)
	}
	//defer ws.Close()
}

func TestWsHandlerMessagesSentAndReceive(t *testing.T) {
	handler := http.HandlerFunc(wsHandler)
	s := httptest.NewServer(handler)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Create correct header
	header := http.Header{}
	header.Set("Content-Type", "text/plain")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, header)
	if err != nil {
		t.Fatalf("Could not open a ws connection on %s %v", u, err)
	}
	//defer ws.Close()

	tables := []struct{ input, output string }{{"ima", "ima"}, {"?sa?", "!sa!"}}
	for _, table := range tables {
		// Send message to server
		err := ws.WriteMessage(1, []byte(table.input))
		if err != nil {
			t.Fatalf("Could not send message over ws %v", err)
		}

		// Receive message from server
		_, p, err := ws.ReadMessage()
		if err != nil {
			t.Fatalf("Could not read message %v", err)
		}

		// Assert message from server
		assert.Equal(t, string(p), table.output, "Question marks should be replaced with exclamation marks")
	}
}

func TestWaitForRequestsBadMessageType(t *testing.T) {
	handler := http.HandlerFunc(wsHandler)
	s := httptest.NewServer(handler)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Create correct header
	header := http.Header{}
	header.Set("Content-Type", "text/plain")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, header)
	if err != nil {
		t.Fatalf("Could not open a ws connection on %s %v", u, err)
	}
	//defer ws.Close()

	tables := []struct{ input, output string }{{"rty", "rty"}, {"?sa?", "!sa!"}}
	for _, table := range tables {
		// Send message to server
		err := ws.WriteMessage(2, []byte(table.input))
		if err != nil {
			t.Fatalf("Could not send message over ws %v", err)
		}

		// Receive message from server
		_, _, err = ws.ReadMessage()
		if err == nil {
			t.Fatalf("Should be NOT nil because message type not text %v", err)
		}
	}
}

func TestReplaceQuestionMarks(t *testing.T) {
	tables := []struct{ input, output string }{{"imam", "imam"}, {"?sa?", "!sa!"}}

	for _, table := range tables {
		result := replaceQuestionMarks(table.input)
		assert.Equal(t, result, table.output, "Question marks should be replaced with exclamation marks")
	}
}

func TestContains(t *testing.T) {
	arrays := [][]string{{"a", "b", "c"}, {"d", "e"}}
	strs := []string{"a", "b"}
	results := []bool{true, false}

	for i, arr := range arrays {
		assert.Equal(t, results[i], contains(arr, strs[i]), "String should contain or not in an array")
	}
}
