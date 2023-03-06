package handler

import (
	config "assignment_1/config"
	"net/http"
)

func EmptyHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "No functionality on root level. Please use paths "+config.UNI_INFO_PATH+" or "+config.COUNTRIES_PATH+
		" or "+config.DIAGNOSTICS_PATH+".", http.StatusOK)
}
