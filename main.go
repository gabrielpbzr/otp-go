package main

import (
	"log"
	"net/http"

	"github.com/gabrielpbzr/otp-go/web"
)

func main() {
	mux := http.NewServeMux()

	web.SetWebHandlers(mux)

	msg := "Server is up and running on TCP port 3000"
	log.Println(msg)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Panicln(err)
	}
}
