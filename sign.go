package sign

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"strings"
)

// Sign will sign the request with appKey and appKeySecret.
func Sign(req *http.Request, appKey, appKeySecret string) error {
	req.Header.Set(HTTPHeaderCATimestamp, CurrentTimeMillis())
	req.Header.Set(HTTPHeaderCANonce, UUID4())
	req.Header.Set(HTTPHeaderCAKey, appKey)

	if req.Header.Get(HTTPHeaderAccept) == "" {
		req.Header.Set(HTTPHeaderAccept, defaultAccept)
	}

	if req.Header.Get(HTTPHeaderDate) == "" {
		req.Header.Set(HTTPHeaderDate, CurrentGMTDate())
	}

	if req.Header.Get(HTTPHeaderUserAgent) == "" {
		req.Header.Set(HTTPHeaderUserAgent, defaultUserAgent)
	}

	if req.Body != nil && req.Header.Get(HTTPHeaderContentType) == HTTPContentTypeStream {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		req.Header.Set(HTTPHeaderContentMD5, MD5(b))
	}

	stringToSign, err := buildStringToSign(req)
	if err != nil {
		return err
	}

	// log.Println("\n" + stringToSign + "\n")
	req.Header.Set(HTTPHeaderCASignature, HmacSHA256([]byte(stringToSign), []byte(appKeySecret)))
	return nil
}

func buildStringToSign(req *http.Request) (string, error) {
	s := ""
	s += strings.ToUpper(req.Method) + defaultLF

	s += req.Header.Get(HTTPHeaderAccept) + defaultLF
	s += req.Header.Get(HTTPHeaderContentMD5) + defaultLF
	s += req.Header.Get(HTTPHeaderContentType) + defaultLF
	s += req.Header.Get(HTTPHeaderDate) + defaultLF

	s += buildHeaderStringToSign(req)

	paramStr, err := buildParamStringToSign(req)
	if err != nil {
		return "", err
	}
	s += paramStr

	return s, nil
}

func buildHeaderStringToSign(req *http.Request) string {
	headerKeys := getSortKeys(req.Header)
	headerCAs := make([]string, 0)
	headerCAKeys := make([]string, 0)

	for _, key := range headerKeys {
		if strings.HasPrefix(http.CanonicalHeaderKey(key), HTTPHeaderCAPrefix) {
			headerCAKeys = append(headerCAKeys, key)
			headerCAs = append(headerCAs, key+":"+req.Header.Get(key)+defaultLF)
		}
	}

	req.Header.Set(HTTPHeaderCASignatureHeaders, strings.Join(headerCAKeys, ","))
	return strings.Join(headerCAs, "")
}

func buildParamStringToSign(req *http.Request) (string, error) {
	var err error
	reqClone := req.Clone(context.Background())
	if req.Body != nil {
		reqClone.Body, err = req.GetBody()
		if err != nil {
			return "", err
		}
	}

	err = reqClone.ParseForm()
	if err != nil {
		return "", err
	}

	paramKeys := getSortKeys(reqClone.Form)
	paramList := make([]string, 0)
	for _, key := range paramKeys {
		value := reqClone.Form.Get(key)
		if value == "" {
			paramList = append(paramList, key)
		} else {
			paramList = append(paramList, key+"="+value)
		}
	}

	params := strings.Join(paramList, "&")
	if params != "" {
		params = "?" + params
	}
	return req.URL.Path + params, nil
}
