
コマンド
```
go test -v .
go test -cover .
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
go test -coverprofile=coverage.out && go tool cover -html=coverage.out

go test -run Test_isPrime
go test -run Test_isPrime -v
```
