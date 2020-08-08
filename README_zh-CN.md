# 阿里云 API 网关请求签名算法的 Go 实现

*[English](README.md) ∙ [简体中文](README_zh-CN.md)*

[![Github License](https://img.shields.io/github/license/sliveryou/aliyun-api-gateway-sign.svg?style=flat)](https://github.com/sliveryou/aliyun-api-gateway-sign/blob/master/LICENSE)
[![Go Doc](https://godoc.org/github.com/sliveryou/aliyun-api-gateway-sign?status.svg)](https://pkg.go.dev/github.com/sliveryou/aliyun-api-gateway-sign)
[![Go Report](https://goreportcard.com/badge/github.com/sliveryou/aliyun-api-gateway-sign)](https://goreportcard.com/report/github.com/sliveryou/aliyun-api-gateway-sign)
[![Github Latest Release](https://img.shields.io/github/release/sliveryou/aliyun-api-gateway-sign.svg?style=flat)](https://github.com/sliveryou/aliyun-api-gateway-sign/releases/latest)
[![Github Latest Tag](https://img.shields.io/github/tag/sliveryou/aliyun-api-gateway-sign.svg?style=flat)](https://github.com/sliveryou/aliyun-api-gateway-sign/tags)
[![Github Stars](https://img.shields.io/github/stars/sliveryou/aliyun-api-gateway-sign.svg?style=flat)](https://github.com/sliveryou/aliyun-api-gateway-sign/stargazers)

详细的签名算法，请查看：[客户端签名说明文档](https://help.aliyun.com/document_detail/29475.html)

## 安装

使用如下命令下载并安装包：

```sh
$ go get github.com/sliveryou/aliyun-api-gateway-sign
```

## 使用示例

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

	// 先创建一个 HTTP 请求对象
	req, err := http.NewRequest(sign.HTTPMethodPost, url, strings.NewReader(body))
	if err != nil {
		// 错误逻辑处理
		panic(err)
	}
	// 设置请求的请求头
	req.Header.Set(sign.HTTPHeaderAccept, sign.HTTPContentTypeJson)
	req.Header.Set(sign.HTTPHeaderContentType, sign.HTTPContentTypeJson)

	// 对请求使用签名函数进行签名
	if err := sign.Sign(req, appKey, appKeySecret); err != nil {
		panic(err)
	}

	// 打印签名后请求的具体内容
	dumpReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		panic(err)
	}
	log.Println("\n" + string(dumpReq))

	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 处理返回的响应内容
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println("\n"+string(content), resp.StatusCode, resp.Header.Get("X-Ca-Error-Message"))
}
```
