package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Base64 experiment !!")
	msg := "Hello world !!"
	// encoded message
	encoded := base64.StdEncoding.EncodeToString([]byte(msg))
	fmt.Println(encoded)

	// decoded message
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Decode error=", err)
		return
	}
	fmt.Println("Decoded message =", string(decoded))

}
