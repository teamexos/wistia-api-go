# wistia-api-go [![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/teamexos/wistia-api-go%2Ftest?type=cf-2&key=eyJhbGciOiJIUzI1NiJ9.NWUzMjVkOWM5M2VlYzM4YjIzZGU4ZGE1.EAROfiOaFnHfuH1BC0B545cqClv1Hwyu87_GpJhVhVo)]( https://g.codefresh.io/pipelines/edit/new/builds?id=5f656fd759f8fff6378ef74d&pipeline=test&projects=wistia-api-go&projectId=5f656fbb59f8fff5ed8ef74c)

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