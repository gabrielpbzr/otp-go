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
}

func generateOtpHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("ERROR: Error parsing the request " + err.Error())
		return
	}

	secret := r.PostForm.Get("secret")
	if secret == "" {
		log.Println("Request doesn't sent any secret. Abort OTP generation")
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
	secret := r.URL.Query().Get("secret")
	log.Printf("Secret %s", secret)
	fmt.Fprintln(w, encoding.EncodeSecret(secret))
}
