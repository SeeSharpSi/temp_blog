package main

import (
	routes "blog/main/routing"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
    port := flag.Int("port", 4000, "port for localhost")
    flag.Parse()
    port_value := *port
	ip := "127.0.0.1:" + strconv.Itoa(port_value)
	wg.Add(1)
	go runServer(ip, &wg)
	fmt.Printf("Server running on %s\n", ip)
	wg.Wait()
}

func runServer(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	mux := routes.StartHandlers()
	server := http.Server{
		Addr:    port,
		Handler: mux,
	}
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
