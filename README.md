# Simple implementation of TOTP for 2FA in Go
This implementation uses SHA-1 as hash algorithm.


## Build
```shell
go build -o totp main.go
```

### Usage
```shell
./totp --secret-key some-secret-key
```
The secret key must be encoded in base32.

## Credits
Based on this blog [https://drewdevault.com/2022/10/18/TOTP-is-easy.html](https://drewdevault.com/2022/10/18/TOTP-is-easy.html).
