package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Msg type definition
type Msg struct {
	User string `json: user`
	Mesg string `json: mesg`
}

var remoteIP string
var myName string

func main_() {

	remoteIP = os.Args[1]
	myName = os.Args[2]
	go chatClient()
	mux := http.NewServeMux()
	mux.HandleFunc("/", chatServer)
	http.ListenAndServe(":8081", mux)

}

// Monitor service
func chatServer(w http.ResponseWriter, r *http.Request) {
	// endpoint listening for status from target, returns the current status of task in json

	decoder := json.NewDecoder(r.Body)
	var s Msg
	err := decoder.Decode(&s)
	if err != nil {
		panic(err)
	}
	fmt.Println(s.User + ": " + s.Mesg)
	fmt.Fprintf(w, "received")

}

func chatClient() {

	var m Msg
	m.User = myName

	reader := bufio.NewReader(os.Stdin)

	for {
		_msg, err := reader.ReadString('\n')
		m.Mesg = _msg
		payloadJSON, err := json.Marshal(m)
		if err != nil {
			fmt.Println("ERROR:failed to parse payload json")
			panic(err)
		}

		url := "http://" + remoteIP + ":8081/"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadJSON))
		if err != nil {
			fmt.Println("ERROR:failed to create request")
			panic(err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("ERROR:request failed")
			panic(err)
		}
		resp.Body.Close()
	}

}
