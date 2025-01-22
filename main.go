package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("need port number\n")
		os.Exit(1)
	}

	p := os.Args[1]
	l, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("failed to listen port %s: %v", p, err)
	}

	err = run(context.Background(), l)
	if err != nil {
		log.Printf("failed to terminate server: %+v", err)
	}
}

func run(ctx context.Context, l net.Listener) (err error) {
	s := &http.Server{
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
			},
		),
	}
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(
		func() error {
			err := s.Serve(l)
			if err != nil && err != http.ErrServerClosed {
				log.Printf("failed to close: %+v", err)
				return err
			}

			return nil
		},
	)

	<-ctx.Done()
	err = s.Shutdown(context.Background())
	if err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}

	return eg.Wait()
}
