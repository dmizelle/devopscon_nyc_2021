package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// make our logger a little human friendly
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	// create a channel so that we can tell the main thread we're ready to stop
	done := make(chan bool, 1)

	// create a channel to send OS Signals to the http server
	exit := make(chan os.Signal, 1)

	// setup a handler to send OS Signals to the channel
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	// create a basic http Hello, World! Server
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Info().Msg("received request, sleeping")

		time.Sleep(10 * time.Second)

		log.Info().Msg("sleep done")

		fmt.Fprintf(w, "Hello DevOpsCon NYC 2021!")
	})
	server := &http.Server{
		Addr:    ":8888",
		Handler: mux,
	}

	// background a thread to wait for signals and shutdown the http listener
	go func(server *http.Server, done chan bool, exit <-chan os.Signal) {
		// wait for us to receive an OS signal
		sig := <-exit

		log.Info().Str("signal", sig.String()).Msg("received signal")

		// create a "deadline" for waiting the server to stop gracefully
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// turn off keepalives, essentially terminating idle connections
		server.SetKeepAlivesEnabled(false)

		// tell the http listener to stop receiving new connections
		// and handle any leftover requests
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal().Err(err).Msg("unable to gracefully shutdown")
		}

		log.Info().Msg("gracefully shutdown")

		// tell the main thread we've shut down
		close(done)

	}(server, done, exit)

	log.Info().Msg("starting server")

	// start the http server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal().Err(err).Msg("unable to start server")
	}

	// block waiting for the graceful shutdown to occur
	<-done

	log.Info().Msg("stopped server")
}
