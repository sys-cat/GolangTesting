# GolangTesting

something researches in golang.

## Connection Encrypto

### Use HMAC using sha512

```bash
go test -bench . -benchmem | prettybench
testing: warning: no tests to run
--- BENCH: BenchmarkRun-8
    hmac_test.go:9: LeebJMITvGM88uFxX4TzzfbM9HRIbcH5zsOKviznr97Pv+wmyd5sIyMKiqoa1dwFqIRrhcnykhxByV7TCjCw/g==
    hmac_test.go:9: LeebJMITvGM88uFxX4TzzfbM9HRIbcH5zsOKviznr97Pv+wmyd5sIyMKiqoa1dwFqIRrhcnykhxByV7TCjCw/g==
    hmac_test.go:9: LeebJMITvGM88uFxX4TzzfbM9HRIbcH5zsOKviznr97Pv+wmyd5sIyMKiqoa1dwFqIRrhcnykhxByV7TCjCw/g==
    hmac_test.go:9: LeebJMITvGM88uFxX4TzzfbM9HRIbcH5zsOKviznr97Pv+wmyd5sIyMKiqoa1dwFqIRrhcnykhxByV7TCjCw/g==
    hmac_test.go:9: LeebJMITvGM88uFxX4TzzfbM9HRIbcH5zsOKviznr97Pv+wmyd5sIyMKiqoa1dwFqIRrhcnykhxByV7TCjCw/g==
    hmac_test.go:9: LeebJMITvGM88uFxX4TzzfbM9HRIbcH5zsOKviznr97Pv+wmyd5sIyMKiqoa1dwFqIRrhcnykhxByV7TCjCw/g==
PASS
benchmark                      iter    time/iter   bytes alloc        allocs
---------                      ----    ---------   -----------        ------
BenchmarkRun-8           2000000000   0.00 ns/op        0 B/op   0 allocs/op
BenchmarkRun100Times-8   2000000000   0.00 ns/op        0 B/op   0 allocs/op
ok      github.com/sys-cat/GolangTesting/hmac    0.010s
```

### Use JWT using HS512

```bash
go test -bench . -benchmem | prettybench
testing: warning: no tests to run
--- BENCH: BenchmarkRun-8
    jwt_test.go:6: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoxLCJQYXNzIjoicGFzc3dvcmQiLCJQYXlsb2FkIjoic29tZXRoaW5nIHBheWxvYWRzIn0.3gkqwMV0cGQ_y6IBXx2XKFij3-08sDdImguRMGUEBQDjfaWFAWLdrq-f5DFoUsNRBsHSJI2zSFVcjwsAb0GW5w
    jwt_test.go:6: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoxLCJQYXNzIjoicGFzc3dvcmQiLCJQYXlsb2FkIjoic29tZXRoaW5nIHBheWxvYWRzIn0.3gkqwMV0cGQ_y6IBXx2XKFij3-08sDdImguRMGUEBQDjfaWFAWLdrq-f5DFoUsNRBsHSJI2zSFVcjwsAb0GW5w
    jwt_test.go:6: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoxLCJQYXNzIjoicGFzc3dvcmQiLCJQYXlsb2FkIjoic29tZXRoaW5nIHBheWxvYWRzIn0.3gkqwMV0cGQ_y6IBXx2XKFij3-08sDdImguRMGUEBQDjfaWFAWLdrq-f5DFoUsNRBsHSJI2zSFVcjwsAb0GW5w
    jwt_test.go:6: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoxLCJQYXNzIjoicGFzc3dvcmQiLCJQYXlsb2FkIjoic29tZXRoaW5nIHBheWxvYWRzIn0.3gkqwMV0cGQ_y6IBXx2XKFij3-08sDdImguRMGUEBQDjfaWFAWLdrq-f5DFoUsNRBsHSJI2zSFVcjwsAb0GW5w
    jwt_test.go:6: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoxLCJQYXNzIjoicGFzc3dvcmQiLCJQYXlsb2FkIjoic29tZXRoaW5nIHBheWxvYWRzIn0.3gkqwMV0cGQ_y6IBXx2XKFij3-08sDdImguRMGUEBQDjfaWFAWLdrq-f5DFoUsNRBsHSJI2zSFVcjwsAb0GW5w
    jwt_test.go:6: eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjoxLCJQYXNzIjoicGFzc3dvcmQiLCJQYXlsb2FkIjoic29tZXRoaW5nIHBheWxvYWRzIn0.3gkqwMV0cGQ_y6IBXx2XKFij3-08sDdImguRMGUEBQDjfaWFAWLdrq-f5DFoUsNRBsHSJI2zSFVcjwsAb0GW5w
PASS
benchmark                      iter    time/iter   bytes alloc        allocs
---------                      ----    ---------   -----------        ------
BenchmarkRun-8           2000000000   0.00 ns/op        0 B/op   0 allocs/op
BenchmarkRun100Times-8   2000000000   0.00 ns/op        0 B/op   0 allocs/op
ok      github.com/sys-cat/GolangTesting/jwt    0.075s
```