package main

import (
	"encoding/base32"
	"flag"
	"fmt"
	"time"

	"github.com/tupyy/totp/totp"
)

func main() {
	var secret string
	flag.StringVar(&secret, "secret-key", "", "Secret key encoded base32")
	flag.Parse()

	if secret == "" {
		fmt.Println("secret key is missing")
		return
	}

	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		panic(err)
	}
	instant := time.Now().UTC()

	code := totp.Totp(instant, key)

	fmt.Println(code)
}
