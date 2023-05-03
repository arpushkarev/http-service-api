package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Key   string
	Value int64
}

func main() {
	fmt.Println("Введите запрос:")
	sbuf := bufio.NewReader(os.Stdin)
	s, err := sbuf.ReadBytes('\n')
	var data Person

	err = json.Unmarshal(s, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data.Value + 1)
}
