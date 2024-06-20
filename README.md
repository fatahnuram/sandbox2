# sandbox2

Following `Build SaaS apps in Go` by `Dominic St-Pierre`.

## How to run?

Run:

```bash
go run cmd/main.go
```

Then try to hit your web server:

```bash
curl -v localhost:8080
```

## How to build?

Run:

```bash
go build -o svc ./cmd/main.go
```

Then run your app:

```bash
./svc
```

Then try to hit your web server:

```bash
curl -v localhost:8080
```
