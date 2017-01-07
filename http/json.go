package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//ResponseJSON middleware to write json to http
func ResponseJSON(w http.ResponseWriter, data interface{}, code int, callback string) {
	b, err := json.Marshal(data)

	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
		return
	}

	if callback != "" {
		w.Header().Set("Content-Type", "text/javascript")
		fmt.Fprintf(w, "%s(%s)", callback, b)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}

//ResponseJSONCode middle to write http header code to header
func ResponseJSONCode(w http.ResponseWriter, message string, code int) {
	type response struct {
		M string `json:"message"`
	}

	b, err := json.Marshal(response{M: message})
	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(b)
}
