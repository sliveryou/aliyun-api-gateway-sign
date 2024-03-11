# Aliyun api gateway request signature algorithm implemented by go

*[English](README.md) ∙ [简体中文](README_zh-CN.md)*

[![Github License](https://img.shields.io/github/license/sliveryou/aliyun-api-gateway-sign.svg?style=flat)](https://github.com/sliveryou/aliyun-api-gateway-sign/blob/master/LICENSE)
[![Go Doc](https://godoc.org/github.com/sliveryou/aliyun-api-gateway-sign?status.svg)](https://pkg.go.dev/github.com/sliveryou/aliyun-api-gateway-sign)
[![Go Report](https://goreportcard.com/badge/github.com/sliveryou/aliyun-api-gateway-sign)](https://goreportcard.com/report/github.com/sliveryou/aliyun-api-gateway-sign)
[![Github Latest Release](https://img.shields.io/github/release/sliveryou/aliyun-api-gateway-sign.svg?style=flat)](https://github.com/sliveryou/aliyun-api-gateway-sign/releases/latest)
[![Github Latest Tag](https://img.shields.io/github/tag/sliveryou/aliyun-api-gateway-sign.svg?style=flat)](https://github.com/sliveryou/aliyun-api-gateway-sign/tags)
[![Github Stars](https://img.shields.io/github/stars/sliveryou/aliyun-api-gateway-sign.svg?style=flat)](https://github.com/sliveryou/aliyun-api-gateway-sign/stargazers)

For signature methods, see: [using digital signature authentication to call api](https://help.aliyun.com/document_detail/29475.html)

## Installation

Download package by using:

```sh
$ go get github.com/sliveryou/aliyun-api-gateway-sign
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
		panic(err)
	}

	// Show the dump request.
	dumpReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	}
	log.Println("\n" + string(dumpReq))

	// Do the request.
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Handle response.
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("\n"+string(content), resp.StatusCode, resp.Header.Get("X-Ca-Error-Message"))
}
```
