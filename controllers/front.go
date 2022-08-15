package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc) // /users/5
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {

	enc := json.NewEncoder(w)
	enc.Encode(data)

}
