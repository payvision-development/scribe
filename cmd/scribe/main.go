package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"

	"github.com/gorilla/mux"
	"github.com/payvision-development/scribe"
	"github.com/payvision-development/scribe/freshservice"
	"github.com/payvision-development/scribe/health"
	"github.com/payvision-development/scribe/release"
	"github.com/payvision-development/scribe/vss"
)

// Specification struct
type Specification struct {
	FreshserviceURL    string `required:"true" split_words:"true"`
	FreshserviceApikey string `required:"true" split_words:"true"`
}

var env Specification
var events = make(chan *vss.Event, 10)

func main() {
	fmt.Println("Scribe is alive")

	err := envconfig.Process("scribe", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	go eventRouter()

	router := mux.NewRouter()
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/vss-release", vssRelease).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func status(rw http.ResponseWriter, req *http.Request) {
	status := health.Status{
		Service:     "Scribe",
		Description: "VSTS Release event integration with Freshservice",
		Status:      "OK",
		Version:     "0.0.0",
	}

	client := freshservice.NewClient(env.FreshserviceURL, env.FreshserviceApikey)

	err := client.CheckEndpoint()
	if err != nil {
		status.Status = "KO: Failed to connect with Freshservice API endpoint"
	}

	json.NewEncoder(rw).Encode(status)
}

func vssRelease(rw http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	event, err := scribe.Parser(b)
	if err != nil {
		fmt.Println(err)
		return
	}

	events <- event
}

func eventRouter() {
	m := make(map[uint32]chan *vss.Event)

	for event := range events {
		deploy, ok := m[event.ReleaseTrackingCode]

		if !ok {
			fmt.Printf("[Release: %v] Creating new session...\n", event.ReleaseTrackingCode)

			deploy = make(chan *vss.Event)
			m[event.ReleaseTrackingCode] = deploy

			client := freshservice.NewClient(env.FreshserviceURL, env.FreshserviceApikey)
			changer := release.FreshserviceChanger{Client: client}

			go scribe.Session(event.ReleaseTrackingCode, deploy, &changer)
		}

		deploy <- event
	}
}
