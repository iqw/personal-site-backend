package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "hello\n")
	if err != nil {
		log.Errorf("Unable to write message to an output")
	}
}

func world(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "world\n")
	if err != nil {
		log.Errorf("Unable to write message to an output")
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			_, err := fmt.Fprintf(w, "%v: %v\n", name, h)
			if err != nil {
				log.Errorf("Unable to write message to an output")
			}
		}
	}
}

func main() {
	viper.AutomaticEnv()
	viper.SetDefault("PORT", 8080)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	http.HandleFunc("/headers", headers)

	log.Infof("Server started at port %s...", viper.GetString("PORT"))
	err := http.ListenAndServe(":"+viper.GetString("PORT"), nil)
	if err != nil {
		panic(fmt.Sprintf("Unable to start a web server on port %s", viper.GetString("PORT")))
	}
}
