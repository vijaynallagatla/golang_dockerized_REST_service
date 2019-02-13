package features

import (
	Model "AuditLog_service/model"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func logInfo(w http.ResponseWriter, r *http.Request) {
	var l Model.Logger
	if err := json.NewDecoder(r.Body).Decode(&l); err != nil {
		log.Println("Error: ", err)
	}
	defer r.Body.Close()
	if err := json.NewEncoder(w).Encode("Success"); err != nil {
		log.WithFields(log.Fields{
			"Error": "Error in sending response",
		})
	}
	log.WithFields(log.Fields{
		"AUDIT_LOG": l.AuditLog,
		"UUID":      l.UUID,
	}).Info(l.Message)
}
