package handlers

import "net/http"


func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, x-requested-with")
	w.Header().Set("Access-Control-Max-Age", "86400")
}