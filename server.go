package main

import (
	routes "blog/main/routing"
	"errors"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	port := "127.0.0.1:42069"
	wg.Add(1)
	go runServer(port, &wg)
	fmt.Printf("Server running on %s\n", port)
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
