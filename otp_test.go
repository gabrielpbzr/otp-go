package main

import "testing"

func TestShouldGenerateOTPForAGivenSecret(t *testing.T) {
	secret := []byte("mysecret")
	otp, err := GenerateOTP(secret)

	if err != nil {
		t.Error(err)
	}

	if otp == "" {
		t.Error("No OTP generated")
	}
}
