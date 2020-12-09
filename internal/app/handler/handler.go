package handler

import (
	"fmt"
	"net/http"

	"github.com/triardn/golang-base-project/internal/app/commons"
	"github.com/triardn/golang-base-project/internal/app/service"
)

// HandlerOption option for handler, including all service
type HandlerOption struct {
	commons.Options
	*service.Services
}

type HttpHandler struct {
	H func(w http.ResponseWriter, r *http.Request) error
}

func (h HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(w, r)
	if err != nil {
		// handle returned error here.
		fmt.Println(err)

		w.WriteHeader(503)
		w.Write([]byte("Bad Gateway"))
	}
}
