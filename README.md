# GolangTesting
something researches in golang.

## Use HMAC using sha512

```
go test -bench . -benchmem                                                                                                                                                                           ⮂
testing: warning: no tests to run
BenchmarkRun-8   	2000000000	         0.00 ns/op	       0 B/op	       0 allocs/op
--- BENCH: BenchmarkRun-8
	hmac_test.go:8: nDKKwCF36GeBHKQR%2BnQuoxQrCkO7zjwD%2Fpi%2FgYYSB5SgXlX8b9IMHXOnL091Hzn017jDZma1HdW4ebdiTKOb7Q%3D%3D
	hmac_test.go:8: nDKKwCF36GeBHKQR%2BnQuoxQrCkO7zjwD%2Fpi%2FgYYSB5SgXlX8b9IMHXOnL091Hzn017jDZma1HdW4ebdiTKOb7Q%3D%3D
	hmac_test.go:8: FUr2dGEGP4wsXadE%2FvdotYhv1CXaYcTaolWp%2BdWwsY%2FHqvSxe%2F51YqqIQ4xhTS9bLY6HVfLXDrEbmZ8WqWrFVA%3D%3D
	hmac_test.go:8: FUr2dGEGP4wsXadE%2FvdotYhv1CXaYcTaolWp%2BdWwsY%2FHqvSxe%2F51YqqIQ4xhTS9bLY6HVfLXDrEbmZ8WqWrFVA%3D%3D
	hmac_test.go:8: nDKKwCF36GeBHKQR%2BnQuoxQrCkO7zjwD%2Fpi%2FgYYSB5SgXlX8b9IMHXOnL091Hzn017jDZma1HdW4ebdiTKOb7Q%3D%3D
	hmac_test.go:8: FUr2dGEGP4wsXadE%2FvdotYhv1CXaYcTaolWp%2BdWwsY%2FHqvSxe%2F51YqqIQ4xhTS9bLY6HVfLXDrEbmZ8WqWrFVA%3D%3D
PASS
ok  	github.com/sys-cat/GolangTesting/hmac	0.007s
```