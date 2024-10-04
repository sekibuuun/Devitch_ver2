package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sekibuuun/Devitch_ver2/backend/controllers/services"
)

type HealthCheckController struct {
	service services.HealthCheckServicer
}

func NewHealthCheckController(s services.HealthCheckServicer) *HealthCheckController {
	return &HealthCheckController{service: s}
}

func (c *HealthCheckController) HelloHandler(w http.ResponseWriter, r *http.Request) {
	hellos, err := c.service.GetHelloService()
	if err != nil {
		log.Println("Could not get hellos")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(&hellos); err != nil {
		log.Println("Could not encode hellos to JSON")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Println(buf.String())

	_, err = fmt.Fprint(w, buf.String())
	if err != nil {
		log.Println("Error writing response:", err)
	}
}
