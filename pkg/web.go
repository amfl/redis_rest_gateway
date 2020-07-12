package redis_rest_gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

// Payload represents a message published back to Redis in JSON
type Payload struct {
	Header http.Header
	Data   map[string](interface{})
}

// Gateway is the webserver itself
type Gateway struct {
	Client  *redis.Client
	Context context.Context
}

func (gw *Gateway) homeLink(w http.ResponseWriter, r *http.Request) {
	// Get the channel we're publishing to from the URL vars
	vars := mux.Vars(r)
	channel := vars["channel"]

	// Create a new payload object
	payload := Payload{
		Header: r.Header,
	}

	// Slurp the JSON from the original request and put it in the payload.
	// We could omit this step and copy the data directly into the payload as a
	// string instead, but my input data is already JSON, and embedding JSON
	// strings inside JSON makes me sad.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading web request: ", err)
	}
	err = json.Unmarshal(body, &payload.Data)
	if err != nil {
		fmt.Printf("Err receiving: %s", err)
		fmt.Fprintf(w, err.Error())
		return
	}
	defer r.Body.Close()

	// Convert payload to JSON so we can send it over the wire on redis
	jsonBytes, _ := json.Marshal(payload)
	jsonString := string(jsonBytes)

	// Send to redis
	fmt.Println("Sending: ", jsonString)
	err = gw.Client.Publish(gw.Context, channel, jsonString).Err()
	if err != nil {
		fmt.Printf("Err sending: %s", err)
		fmt.Fprintf(w, err.Error())
	}

	// Write back a success message to the web client
	fmt.Fprintf(w, "OK")
}

// Listen will start the web gateway listening for incoming requests.
func (gw *Gateway) Listen(listenInterface string) {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/channel/{channel}", gw.homeLink)
	log.Fatal(http.ListenAndServe(listenInterface, router))
}
