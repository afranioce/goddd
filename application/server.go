package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/afranioce/goddd/config"
)

// RunHTTPServer provide run http or https protocol.
func RunHTTPServer() {
	srv := &http.Server{
		Addr:    config.MustEnv("LISTEN_ADDR"),
		Handler: routerEngine(),
	}

	go func() {
		sslEnabled, err := strconv.ParseBool(config.MustEnv("SSL_ENABLED"))
		// service connections
		if err == nil {
			if sslEnabled == true {
				err = srv.ListenAndServeTLS(os.Getenv("SSL_CERT_FILE"), os.Getenv("SSL_CERT_FILE"))
			} else {
				err = srv.ListenAndServe()
			}
		}

		if err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
