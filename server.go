package main

import (
	sql_funcs "blog/main/db"
	routes "blog/main/routing"
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/yuin/goldmark"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var wg sync.WaitGroup
	port := flag.Int("port", 9779, "port for localhost")
	add_post := flag.Bool("ap", false, "add a post")
	flag.Parse()
	if *add_post {
		adding_post()
	} else {
		ip := "127.0.0.1:" + strconv.Itoa(*port)
		wg.Add(1)
		go runServer(ip, &wg)
		fmt.Printf("Server running on %s\n", ip)
		wg.Wait()
	}
}

func runServer(port string, wg *sync.WaitGroup) {
	defer wg.Done()
	mux := routes.StartHandlers()
	server := http.Server{
		Addr:    port,
		Handler: mux,
	}
	err := server.ListenAndServe()
	defer server.Close()
	defer fmt.Println("Closed server")
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func adding_post() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Title: ")
	title, err := reader.ReadString('\n')
	check(err)
	title = strings.Replace(title, "\n", "", -1)
	fmt.Println("Teaser: ")
	teaser, err := reader.ReadString('\n')
	check(err)
	teaser = strings.Replace(teaser, "\n", "", -1)
	fmt.Println("Content (relative file location): ")
	//content_file_location, err := reader.ReadString('\n')
    file, err := os.ReadFile("temp.txt")
    check(err)
    var buf bytes.Buffer
    err = goldmark.Convert(file, &buf)
    check(err)
    content := buf.String()
	sql_funcs.Add_Post(title, teaser, content)
}
