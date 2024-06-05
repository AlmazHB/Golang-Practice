package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type MyHandlaer struct {
	Router *mux.Router
	Http   *http.Server
}

func ServerProvider() *MyHandlaer {
	mH := &MyHandlaer{}
	mH.Router = mux.NewRouter()
	mH.mapRoutes()

	mH.Http = &http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: mH.Router,
	}
	fmt.Println("Trying to create a server")
	return mH
}

func (mH *MyHandlaer) mapRoutes() {
	mH.Router.HandleFunc("/soltan", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is a response from the server")
	})
}
func (mH *MyHandlaer) MyServer() error {
	if err := mH.Http.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
