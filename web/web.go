package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabrielpbzr/otp-go/encoding"
	"github.com/gabrielpbzr/otp-go/password"
)

func SetWebHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/", generateOtpHandler)
	mux.HandleFunc("/encode", encodeSecretHandler)
	mux.HandleFunc("/openapi", showOpenAPIDocs)
}

func generateOtpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "POST is the only method allowed for this endpoint")
		return
	}
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error parsing the request: %s", err.Error())
		log.Println("ERROR: Error parsing the request " + err.Error())
		return
	}

	secret := r.PostForm.Get("secret")
	if secret == "" {
		log.Println("Request doesn't sent any secret. Abort OTP generation")
		fmt.Fprintln(w, "Parameter \"secret\" was not provided. Secret is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	otp, err := password.GenerateOTP([]byte(secret))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Error parsing the request " + err.Error())
		return
	}

	fmt.Fprintln(w, otp)
}

func encodeSecretHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "GET is the only method allowed for this endpoint")
		return
	}

	secret := r.URL.Query().Get("secret")
	log.Printf("Secret %s", secret)
	if secret == "" {
		log.Println("Request doesn't sent any secret. Abort OTP generation")
		fmt.Fprintln(w, "Parameter \"secret\" was not provided. Secret is required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, encoding.EncodeSecret(secret))
}

func showOpenAPIDocs(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/openapi.json")
}
