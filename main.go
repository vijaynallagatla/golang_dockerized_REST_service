package main

import (
	"net/http"
	"os"
	"time"

	applicationRoute "AuditLog_service/router"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func logMessageEveryTenSecond() {
	for {
		log.WithFields(log.Fields{
			"AUDIT_LOG": "PNM",
			"UUID":      "abcnsghuerukerdgfg",
		}).Info("Audit Log sample messages")
		time.Sleep(60 * time.Second)
	}
}
func main() {
	go logMessageEveryTenSecond()
	router := httprouter.New()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200", "http://127.0.0.1:4200"},
		AllowCredentials: false,
		Debug:            false,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
	})
	handler := c.Handler(router)
	applicationRoute.ApplicationRoutes(router)
	log.Printf("Server started. \n Listening at 0.0.0.0:4004")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
