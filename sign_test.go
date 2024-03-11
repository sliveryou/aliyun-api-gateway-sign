package sign

import (
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"testing"
)

func TestSignMethodGet(t *testing.T) {
	rawURL := "https://bankcard4c.shumaidata.com/bankcard4c"
	values := make(url.Values)
	values.Add("bankcard", "123456")
	values.Add("idcard", "123456")
	values.Add("mobile", "123456")
	values.Add("name", "测试")
	rawURL += "?" + values.Encode()

	req, err := http.NewRequest(HTTPMethodGet, rawURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = Sign(req, "123456", "123456")
	if err != nil {
		t.Fatal(err)
	}

	dumpReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("\n" + string(dumpReq))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(content), resp.StatusCode, resp.Header.Get("X-Ca-Error-Message"))
}

func TestSignMethodPostForm(t *testing.T) {
	rawURL := "https://bankcard4c.shumaidata.com/bankcard4c"
	values := make(url.Values)
	values.Add("bankcard", "123456")
	values.Add("idcard", "123456")
	values.Add("mobile", "123456")
	values.Add("name", "测试")
	rawURL += "?" + values.Encode()

	req, err := http.NewRequest(HTTPMethodPost, rawURL, strings.NewReader(`a=1&b=2`))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set(HTTPHeaderContentType, HTTPContentTypeForm)
	err = Sign(req, "123456", "123456")
	if err != nil {
		t.Fatal(err)
	}

	dumpReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("\n" + string(dumpReq))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(content), resp.StatusCode, resp.Header.Get("X-Ca-Error-Message"))
}

func TestSignMethodPostJSON(t *testing.T) {
	rawURL := "https://bankcard4c.shumaidata.com/bankcard4c"
	values := make(url.Values)
	values.Add("bankcard", "123456")
	values.Add("idcard", "123456")
	values.Add("mobile", "123456")
	values.Add("name", "测试")
	rawURL += "?" + values.Encode()

	req, err := http.NewRequest(HTTPMethodPost, rawURL, strings.NewReader(`{"a":1,"b":2}`))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set(HTTPHeaderContentType, HTTPContentTypeJson)
	err = Sign(req, "123456", "123456")
	if err != nil {
		t.Fatal(err)
	}

	dumpReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("\n" + string(dumpReq))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(content), resp.StatusCode, resp.Header.Get("X-Ca-Error-Message"))
}
