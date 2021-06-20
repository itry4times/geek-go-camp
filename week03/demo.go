package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Set routing rules
	// http.HandleFunc("/", Tmp)

	//Use the default DefaultServeMux.
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	cancelCtx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(cancelCtx)
	// 模拟创建server
	srv1 := &http.Server{Addr: ":8080"}
	// srv2 := &http.Server{Addr: ":8088"}

	eg.Go(func() error {
		return StartServer(srv1)
	})

	eg.Go(func() error {
		<-ctx.Done()
		log.Println("Http context done.")
		return srv1.Shutdown(ctx)
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh)

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				log.Println("\n Signal context done.")
				return ctx.Err()
			case <-sigCh:
				log.Println("\n Signal recieved, call cancle.")
				cancel()
			}
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("%+v", err)
	}
}

func StartServer(srv *http.Server) error {
	http.HandleFunc("/", Tmp)
	log.Println("Http server starting at ", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func Tmp(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "version 1\n")
}
