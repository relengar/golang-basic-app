package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SendJSON response from data
func SendJSON(data interface{}, w http.ResponseWriter) {
	resp, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Can't parse data to JSON"))
		fmt.Printf("Recipe marshal problem: %s", err)
	}
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.Header().Add("Connection", "keep-alive")
	w.Header().Add("ETag", `W/"1d-4gKrh1OyPNI4mf2K71aE9/Ybb5Y"`)
	w.Write(resp)
}

// ParseID from string
func ParseID(id string, w http.ResponseWriter) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Cannot convert to OID"))
		fmt.Printf("Can't parse string %s to ObjectID: %s", id, err)
	}
	return oid
}

// Throw error response
func Throw(msg string, status int, err error, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
	fmt.Printf("\nError(%s): %s", msg, err)
}
