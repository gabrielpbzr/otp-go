package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"time"
)

func GenerateOTP(secret []byte) (string, error) {
	// Contador Unix time em segundos, dividido pelo intervalo de 30 segundos
	counter := time.Now().Unix() / 30

	// Decodifica a chave de segredo em Base-32
	key := make([]byte, 64)
	_, err := base32.StdEncoding.Decode(key, bytes.ToUpper([]byte(secret)))

	if err != nil {
		return "", err
	}

	// Escreve o contador como 8 bytes, big-endian
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(counter))

	// Calcular o HMAC-SHA1
	mac := hmac.New(sha1.New, key)
	mac.Write(b)
	b = mac.Sum(nil)

	// Pega os últimos 4 bits do  último byte
	offset := b[len(b)-1] & 0xf

	// Lê os 4 bytes do offset como um inteiro de 32 bits
	n := binary.BigEndian.Uint32(b[offset : offset+4])

	// Covert it to decimal
	s := fmt.Sprintf("%06d", int(n&0x7fffffff))

	// Return last 6 digits
	return s[len(s)-6:], nil
}
