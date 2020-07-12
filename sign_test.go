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
	rawurl := "https://bankcard4c.shumaidata.com/bankcard4c"
	values := make(url.Values)
	values.Add("bankcard", "123456")
	values.Add("idcard", "123456")
	values.Add("mobile", "123456")
	values.Add("name", "123456")
	rawurl += "?" + values.Encode()

	req, err := http.NewRequest(HTTPMethodGet, rawurl, nil)
	if err != nil {
		t.Error(err)
		return
	}

	err = Sign(req, "123456", "123456")
	if err != nil {
		t.Error(err)
		return
	}

	dumpReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("\n" + string(dumpReq))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(content), resp.StatusCode, resp.Header.Get("X-Ca-Error-Message"))
}

func TestSignMethodPost(t *testing.T) {
	rawurl := "https://bankcard4c.shumaidata.com/bankcard4c"
	values := make(url.Values)
	values.Add("bankcard", "123456")
	values.Add("idcard", "123456")
	values.Add("mobile", "123456")
	values.Add("name", "123456")
	rawurl += "?" + values.Encode()

	req, err := http.NewRequest(HTTPMethodPost, rawurl, strings.NewReader("a=1&b=2"))
	if err != nil {
		t.Error(err)
		return
	}

	req.Header.Set(HTTPHeaderContentType, HTTPContentTypeForm)
	err = Sign(req, "123456", "123456")
	if err != nil {
		t.Error(err)
		return
	}

	dumpReq, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("\n" + string(dumpReq))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(content), resp.StatusCode, resp.Header.Get("X-Ca-Error-Message"))
}
