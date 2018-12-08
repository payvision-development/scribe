package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/payvision-development/scribe"
	"github.com/payvision-development/scribe/freshservice"
	"github.com/payvision-development/scribe/health"
	"github.com/payvision-development/scribe/vss"
	"github.com/urfave/negroni"
)

// Specification struct
type Specification struct {
	User               string `required:"true"`
	Pass               string `required:"true"`
	FreshserviceURL    string `required:"true" split_words:"true"`
	FreshserviceEmail  string `required:"true" split_words:"true"`
	FreshserviceApikey string `required:"true" split_words:"true"`
	VstsApikey         string `required:"true" split_words:"true"`
}

var env Specification
var events = make(chan *vss.Event, 10)

var date = time.Now()
var count uint32

func main() {
	fmt.Println("Scribe is alive")

	err := envconfig.Process("scribe", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	go eventRouter()

	n := negroni.Classic()
	router := mux.NewRouter()

	router.HandleFunc("/status", status).Methods("GET")

	vssRouter := mux.NewRouter().PathPrefix("/vss").Subrouter().StrictSlash(true)
	vssRouter.HandleFunc("/release", vssRelease).Methods("POST")

	router.PathPrefix("/vss").Handler(negroni.New(
		negroni.HandlerFunc(basicAuth),
		negroni.Wrap(vssRouter),
	))

	n.UseHandler(router)
	n.Run(":8080")
}

func basicAuth(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	user, pass, _ := r.BasicAuth()

	if env.User != user || env.Pass != pass {
		rw.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		http.Error(rw, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	next(rw, r)
}

func status(rw http.ResponseWriter, req *http.Request) {
	status := health.Status{
		Service:     "Scribe",
		Description: "VSTS Release event integration with Freshservice",
		Status:      "OK",
		Version:     "1.1.0",
		Info: health.Info{
			Started: date,
			Events:  count,
		},
	}

	client := freshservice.NewClient(env.FreshserviceURL, env.FreshserviceEmail, env.FreshserviceApikey)

	err := client.CheckEndpoint()
	if err != nil {
		status.Status = "KO: Failed to connect with Freshservice API endpoint"
	}

	json.NewEncoder(rw).Encode(status)
}

func vssRelease(rw http.ResponseWriter, req *http.Request) {
	count++

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
		d, ok := m[event.ReleaseTrackingCode]

		if !ok {
			fmt.Printf("[Release: %v] Creating new session for release %v (%v)\n", event.ReleaseTrackingCode, event.ReleaseName, event.ReleaseID)

			d = make(chan *vss.Event)
			m[event.ReleaseTrackingCode] = d

			var v *vss.TFS
			if event.ServerURL != "" || event.CollectionURL != "" {
				v = vss.NewClient(event.ServerURL, event.CollectionURL, env.VstsApikey)
			}

			fs := freshservice.NewClient(env.FreshserviceURL, env.FreshserviceEmail, env.FreshserviceApikey)

			go scribe.Session(event.ReleaseTrackingCode, d, fs, v)
		}

		d <- event
	}
}
