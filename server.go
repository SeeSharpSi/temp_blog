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
	"os/signal"
	"strconv"
	"strings"

	"github.com/yuin/goldmark"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	port := flag.Int("port", 9779, "port for localhost")
	custom_ip := flag.String("ip", "127.0.0.1", "ip for host")
	add_post := flag.Bool("p", false, "add a post")
	flag.Parse()
	if *add_post {
		adding_post()
	} else {
		ip := *custom_ip + ":" + strconv.Itoa(*port)
		fmt.Printf("Server running on %s\n", ip)
		mux := routes.StartHandlers()
		server := http.Server{
			Addr:    ip,
			Handler: mux,
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			s := <-c
			server.Close()
			fmt.Println(s)
		}()
		err := server.ListenAndServe()
		defer server.Close()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
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
	fmt.Println("Content (relative cli location): ")
	content_file_location, err := reader.ReadString('\n')
	content_file_location = content_file_location[:len(content_file_location)-1]
	file, err := os.ReadFile(content_file_location)
	check(err)
	var buf bytes.Buffer
	err = goldmark.Convert(file, &buf)
	check(err)
	content := buf.String()
	sql_funcs.Add_Post(title, teaser, content)
}
