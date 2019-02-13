package router

import (
	logApi "AuditLog_service/features"

	"github.com/julienschmidt/httprouter"
)

// ApplicationRoutes is to handle all the API Requests
func ApplicationRoutes(router *httprouter.Router) {
	logApi.Routes(router)
}
