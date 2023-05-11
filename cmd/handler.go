package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/arpushkarev/http-service-api/internal/model"
	"github.com/arpushkarev/http-service-api/internal/service"
)

var Info model.Info

func answer(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	ID, err := service.NewService().Create(ctx, Data1)

	Info.Add("id", ID)

	idJson, err := json.MarshalIndent(Info.Id, "", "\t")
	if err != nil {
		log.Println("failed converting data to JSON", err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(idJson)
}
