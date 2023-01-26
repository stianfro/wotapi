package v1

import "net/http"

// Hello is a simple hello world handler
func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello 世界"))
}
