package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func getHandle(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status":"OK"}`
	w.Write([]byte(answer))
}

func main() {
	listen, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal(err)
	}
	for {
		_, err := listen.Accept()
		if err != nil {
			log.Println(err)
			break
		}

	}
}
