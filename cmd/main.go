package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arpushkarev/http-service-api/internal/memory"
)

const port = ":9090"

var Data1 *memory.Person

func main() {

	fmt.Println("Введите запрос:")

	buffer := bufio.NewReader(os.Stdin)
	s, err := buffer.ReadBytes('\n')

	err = json.Unmarshal(s, &Data1)
	if err != nil {
		panic(err)
	}

	srv := http.Server{
		Addr:    port,
		Handler: route(),
	}

	fmt.Printf("Server is running on port%s", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to run HTTP server", err.Error())
	}

}
