package main

import (
	"net/http"
	"time"

	"github.com/ipfans/twirp-demo/internal/greetersvc"
	greeterv1 "github.com/ipfans/twirp-demo/rpc/greeter/v1"
	"github.com/twitchtv/twirp"
)

func main() {
	server := &greetersvc.Server{} // implements `GreeterService` interface
	twirpHandler := greeterv1.NewGreeterServiceServer(
		server,
		twirp.WithServerPathPrefix(""), // Default will be `twirp`
	)

	httpServer := &http.Server{
		Addr:              "127.0.0.1:8080",
		Handler:           twirpHandler,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      60 * time.Second,
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
