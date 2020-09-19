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

## License

MIT License

Copyright (c) 2020 EXOS

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
