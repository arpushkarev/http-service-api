package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/arpushkarev/http-service-api/internal/model"
	"github.com/arpushkarev/http-service-api/internal/service"
)

var id model.Id

func answer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	ID, err := service.NewService().Create(ctx, Data1)

	id.Id = ID

	idStr := fmt.Sprintf("id: %s", id.GetId(id.Id))

	idJson, err := json.MarshalIndent(idStr, "", "\t")
	if err != nil {
		log.Println("failed converting data to JSON", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(idJson)
}
