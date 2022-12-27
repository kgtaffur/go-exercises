package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "JavaTpoint@12345!@#$%^&*()"
	strEncode := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Value to be encoded: " + data)
	fmt.Println("Encoded value: " + strEncode)

	fmt.Println()

	fmt.Println("Value to be decoded: " + strEncode)
	strDecode, _ := base64.StdEncoding.DecodeString(strEncode)
	fmt.Println("Decoded value: " + string(strDecode))

	fmt.Println()

	url := "https://golang.org/ref/spec"

	fmt.Println("URL to be encoded: " + url)
	urlEncode := base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println("Encoded url: " + urlEncode)

	fmt.Println("Value to be decoded: " + urlEncode)
	strDecode2, _ := base64.URLEncoding.DecodeString(urlEncode)

	fmt.Println("Decoded value: " + string(strDecode2))
}
