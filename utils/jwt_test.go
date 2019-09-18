package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	token, _ := GenerateToken("togreat", 30 * 24 * time.Hour)
	fmt.Println(token)
	claim, err := ParseToken(token)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v \n", claim)
}
