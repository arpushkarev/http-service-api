package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/arpushkarev/http-service-api/internal/memory"
	"github.com/arpushkarev/http-service-api/internal/service"
)

type Data1 struct {
	Key   string
	Value int64
}

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

	id, err := service.NewService().Create(ctx, person)

	fmt.Println(id)
}
