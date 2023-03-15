package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"flag"
	"fmt"
	"time"
)

/**
- Divide the current Unix timestamp by 30
- Encode it as a 64-bit big endian integer
- Write the encoded bytes to a SHA-1 HMAC initialized with the TOTP shared key
- Let offs = hmac[-1] & 0xF
- Let hash = decode hmac[offs .. offs + 4] as a 32-bit big-endian integer
- Let code = (hash & 0x7FFFFFFF) % 1000000

Credits: https://drewdevault.com/2022/10/18/TOTP-is-easy.html
**/
func totp(instant time.Time, key []byte) uint32 {
	now := instant.Unix() / 30

	// encode now as big endiant
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(now))

	mac := hmac.New(sha1.New, key)
	mac.Write(buf)
	sum := mac.Sum(nil)

	off := sum[len(sum)-1] & 0xF
	h := sum[off : off+4]
	code := (binary.BigEndian.Uint32(h) & 0x7FFFFFFF) % 1000000
	return code
}

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

	code := totp(instant, key)

	fmt.Println(code)
}
