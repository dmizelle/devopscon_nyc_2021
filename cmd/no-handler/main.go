package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Info().Msg("received request, sleeping")

		time.Sleep(10 * time.Second)

		log.Info().Msg("sleep done")

		fmt.Fprintf(w, "Hello DevOpsCon NYC 2021!")
	})

	log.Info().Msg("server started")

	http.ListenAndServe(":8888", nil)
}
