package httpsrv

import (
	"fmt"
	"log"
	"net/http"
)

type EPHandlerFunc func(w http.ResponseWriter, r *http.Request)

type EPHandler struct {
	URL        string
	HandleFunc EPHandlerFunc
}

type HTTPServer struct {
	Name     string
	Port     int
	Handlers []EPHandler
}

func (h HTTPServer) StartServer() {
	h.AddHandler()
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", h.Port), nil))
}

func (h HTTPServer) AddHandler() {

	for i := 0; i < len(h.Handlers); i++ {
		h := h.Handlers[i]
		http.HandleFunc(h.URL, h.HandleFunc)

	}

}
