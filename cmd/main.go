package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/arpushkarev/http-service-api/internal/memory"
	"github.com/arpushkarev/http-service-api/internal/service"
)

type Data1 struct {
	Key   string
	Value int64
}

var id int64

func main() {

	ctx := context.Background()

	fmt.Println("Введите запрос:")
	buffer := bufio.NewReader(os.Stdin)
	s, err := buffer.ReadBytes('\n')
	var data *Data1

	err = json.Unmarshal(s, &data)
	if err != nil {
		panic(err)
	}

	person := &memory.Person{
		Name: data.Key,
		Age:  data.Value,
	}

	id, err = service.NewService().Create(ctx, person)

	fmt.Println(id)

	http.HandleFunc("/", answer)

	port := ":9090"
	fmt.Printf("Server is running on port%s", port)
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Failed to run HTTP server", err.Error())
	}

}

func answer(w http.ResponseWriter, r *http.Request) {
	idStr := GetId(id)
	idJson, err := json.Marshal(idStr)
	if err != nil {
		log.Println("failed converting data to JSON", err.Error())
	}

	s := string(idJson)

	fmt.Fprintf(w, s)
}

func GetId(id int64) string {
	idStr := strconv.Itoa(int(id + 1))
	return idStr
}
