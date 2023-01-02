package encoding

import (
	"fmt"
	"testing"
)

func TestEncodeSecretString(t *testing.T) {
	secretStr := "mysecretstring"
	encodedSecret := EncodeSecret(secretStr)

	if encodedSecret != "NV4XGZLDOJSXI43UOJUW4ZY=" {
		t.Errorf(fmt.Sprintf("Expected encoded secret to be equal %s", encodedSecret))
	}
}

func TestEncodeEmptyString(t *testing.T) {
	secretStr := ""
	encodedSecret := EncodeSecret(secretStr)

	if encodedSecret != "" {
		t.Errorf("Expected encoded secret to be empty")
	}
}
