package Pinger

import (
	"github.com/gorilla/mux"
	"io"
	"net"
	"net/http"
	"time"
)

type ping struct {
	Router *mux.Router
}
func (sA *ping) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func ListerAndServe(address string) (io.Closer, error) {
	l, err := net.Listen("tcp", address)
	var serveCh chan error
	if err != nil {
		return nil, err
	}
	sA := ping{}
	sA.Router = mux.NewRouter()
	sA.Router.HandleFunc("/ping", sA.ping).Methods("GET")

	go func() {
		serveCh <- http.Serve(l, sA.Router)
	}()

	select {
	case <-time.After(3 * time.Second):
		return l, nil
	case err := <-serveCh:
		return nil, err
	}
}


