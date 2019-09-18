package utils

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("togreat")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(token)
	c, err := ParseToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v", c)
}