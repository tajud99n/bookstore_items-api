package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/tajud99n/bookstore_utils-go/errors"
)

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err errors.RestErr) {
	RespondJSON(w, err.Status(), err)
}
