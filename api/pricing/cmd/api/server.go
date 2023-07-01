package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (app *application) serve() error {

	// Declare out http server settings
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// this will be the channel to receive any erros from the STL Shutdown() function
	shutdownError := make(chan error)

	// have this go routine running in the background
	go func() {
		quit := make(chan os.Signal, 1)

		// here we use STL to listen for SIGINT or SIGTEM singals to the application
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		app.logger.PrintInfo("shutting down server", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// passing in the context we just created, we call Shutdown() on our server.
		// This will initiate a graceful shutdown, and allow any open HTTP connections 5 seconds to complete
		shutdownError <- srv.Shutdown(ctx)
	}()

	app.logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  app.config.env,
	})

	// The reason we are looking for "ErrServerClosed" is that it tells us the graceful shutdown
	// was successfull, as Shutdown() will return an http.ErrServerCloser err.
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	app.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
