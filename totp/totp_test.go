package totp

import (
	"encoding/base32"
	"testing"
	"time"
)

func TestTotp(t *testing.T) {
	secret := "3N2OTFHXKLR2E3WNZSYQ===="
	instant := time.Unix(1650183739, 0).UTC()
	key, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		t.Fatal(err)
	}

	expected := uint32(29283)
	got := Totp(instant, key)
	if got != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}
