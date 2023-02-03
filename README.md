## Test
```bash
go test ./... -v
```

## Cover
```bash
go test -cover ./... -v
```

## Cover Profile
```bash
go test -coverprofile coverage.out ./... -v
```

### Open coverage.out
```bash
go tool cover -html="coverage.out"
```