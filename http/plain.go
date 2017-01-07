package http

import (
	"fmt"
	"net/http"
)

//ResponsePlain middleware to write plain text to http
func ResponsePlain(w http.ResponseWriter, message string, code int, callback string) {
	bytedMessage := []byte(message[:])
	if callback != "" {
		w.Header().Set("Content-Type", "text/javascript")
		fmt.Fprintf(w, "%s(%s)", callback, bytedMessage)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bytedMessage)
}

//ResponsePlainCode middle to write http header code to header
func ResponsePlainCode(w http.ResponseWriter, message string, code int) {
	bytedMessage := []byte(message[:])
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(bytedMessage)
}
