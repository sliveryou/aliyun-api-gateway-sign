# Aliyun api gateway request signature package by go

For signature methods, see: [client signature documentation](https://help.aliyun.com/document_detail/29475.html)

## Installation

```
go get github.com/sliveryou/aliyun-api-gateway-sign
```

## Usage Example

```golang
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"

	sign "github.com/sliveryou/aliyun-api-gateway-sign"
)

func main() {
	var url string = "https://bankcard4c.shumaidata.com/bankcard4c"
	var body string
	var appKey, appKeySecret string

	// Prepare a HTTP request.
	req, err := http.NewRequest(sign.HTTPMethodPost, url, strings.NewReader(body))
	if err != nil {
		// Handle err.
		panic(err)
	}
	// Set the request with headers.
	req.Header.Set(sign.HTTPHeaderAccept, sign.HTTPContentTypeJson)
	req.Header.Set(sign.HTTPHeaderContentType, sign.HTTPContentTypeJson)

	// Sign the request.
	if err := sign.Sign(req, appKey, appKeySecret); err != nil {
		// Handle err.
		panic(err)
	}

	// Show the dump request.
	dumpReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		// Handle err.
		panic(err)
	}
	log.Println("\n" + string(dumpReq))

	// Do the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Handle err.
		panic(err)
	}
	defer resp.Body.Close()

	// Handle response.
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// Handle err.
		panic(err)
	}
	log.Println("\n"+string(content), resp.StatusCode, resp.Header.Get("X-Ca-Error-Message"))
}
```