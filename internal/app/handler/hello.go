package handler

import (
	"encoding/json"
	"net/http"
)

// HelloHandler struct for hello handler
type HelloHandler struct {
	HandlerOption
}

// NewHelloHandler initiate new hello handler
func NewHelloHandler(opt HandlerOption) *HelloHandler {
	return &HelloHandler{
		HandlerOption: opt,
	}
}

// Greet greet person. Introduce yourself
func (h HelloHandler) Greet(w http.ResponseWriter, r *http.Request) (err error) {
	message := h.Services.Hello.Greet()

	response := map[string]interface{}{
		"message": message,
	}

	res, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to unmarshal"))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	return
}
