package http

import (
	"fmt"
	"net/http"
)

//ResponsePlain middleware to write plain text to http
func ResponsePlain(w http.ResponseWriter, message string, code int, callback string) {
	bytedMessage := []byte(message[:])
	if callback != "" {
		fmt.Fprintf(w, "%s(%s)", callback, bytedMessage)
		return
	}

	w.WriteHeader(code)
	w.Write(bytedMessage)
}

//ResponsePlainCode middle to write http header code to header
func ResponsePlainCode(w http.ResponseWriter, message string, code int) {
	bytedMessage := []byte(message[:])

	w.WriteHeader(code)
	w.Write(bytedMessage)
}
