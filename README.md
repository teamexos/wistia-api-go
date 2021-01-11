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

httpClient := wistia.DefaultHTTPClient()
wistiaClient := wistia.NewClient(httpClient, "ACCESS_TOKEN")
```

## Testing

```go
go test -v -tags=unit
```