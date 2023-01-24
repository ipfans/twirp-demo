package main

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/ipfans/twirp-demo/internal/greetersvc"
	greeterv1 "github.com/ipfans/twirp-demo/rpc/greeter/v1"
	"github.com/twitchtv/twirp"
)

func TestServer(t *testing.T) {
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
	go httpServer.ListenAndServe()

	client := greeterv1.NewGreeterServiceProtobufClient("http://127.0.0.1:8080", &http.Client{
		Timeout: 10 * time.Second,
	}, twirp.WithClientPathPrefix(""))
	resp, _ := client.SayHello(context.TODO(), &greeterv1.SayHelloRequest{Name: "Kevin"})
	if !cmp.Equal(resp.Message, "Hello, Kevin") {
		t.Errorf("Unexpected response: %v", resp)
	}
}
