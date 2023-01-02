package main

import "encoding/base32"

func EncodeSecret(secret string) string {
	return base32.StdEncoding.EncodeToString([]byte(secret))
}
