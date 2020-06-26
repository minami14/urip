package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/jessevdk/go-flags"
)

type option struct {
	Addr string `short:"a" long:"addr" default:":80"`
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	addr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err)
		log.Println(r)
		return
	}

	if _, err := w.Write([]byte(addr)); err != nil {
		log.Println(err)
		log.Println(r)
	}

	log.Println(addr)
}

func main() {
	var opt option
	if _, err := flags.Parse(&opt); err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    opt.Addr,
		Handler: new(handler),
	}

	sig := make(chan os.Signal)
	defer close(sig)

	go func() {
		fmt.Printf("Start urip http server. %v\n", opt.Addr)
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
			sig <- os.Interrupt
		}
	}()

	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Printf("\nStopping urip server. %v\n", opt.Addr)
	if err := server.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
	fmt.Printf("Stopped urip server. %v\n", opt.Addr)
}