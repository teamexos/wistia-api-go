# wistia-api-go

[Wistia](https://wistia.com/support/developers) Go client.

## Install

```
go get -u github.com/teamexos/wistia-api-go
```

## Usage

```go
import "github.com/teamexos/wistia-api-go"

// Replace ACCESS_TOKEN with your real access token

restClient := wistia.DefaultHTTPClient()
wistiaClient := wistia.NewClient(restClient, "ACCESS_TOKEN")
```

## Testing

```go
go test -v -tags=unit
```