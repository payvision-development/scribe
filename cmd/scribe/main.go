package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"

	"github.com/gorilla/mux"
	"github.com/payvision-development/scribe"
	"github.com/payvision-development/scribe/freshservice"
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
	router.HandleFunc("/vss-release", vssRelease).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
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
			deploy = make(chan *vss.Event)
			m[event.ReleaseTrackingCode] = deploy

			changer := release.FreshserviceChanger{Client: freshservice.NewClient(env.FreshserviceURL, env.FreshserviceApikey)}

			go scribe.Session(deploy, changer)
		}

		deploy <- event
	}
}
