package sign

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Sign will sign the request with appKey and appKeySecret.
func Sign(req *http.Request, appKey, appKeySecret string) error {
	req.Header.Set(HTTPHeaderCAKey, appKey)
	req.Header.Set(HTTPHeaderCANonce, UUID4())
	req.Header.Set(HTTPHeaderCASignatureMethod, defaultSignMethod)
	req.Header.Set(HTTPHeaderCATimestamp, CurrentTimeMillis())

	if req.Header.Get(HTTPHeaderAccept) == "" {
		req.Header.Set(HTTPHeaderAccept, defaultAccept)
	}

	if req.Header.Get(HTTPHeaderDate) == "" {
		req.Header.Set(HTTPHeaderDate, CurrentGMTDate())
	}

	if req.Header.Get(HTTPHeaderUserAgent) == "" {
		req.Header.Set(HTTPHeaderUserAgent, defaultUserAgent)
	}

	ct := req.Header.Get(HTTPHeaderContentType)
	if req.Body != nil && ct != HTTPContentTypeForm &&
		!strings.HasPrefix(ct, HTTPContentTypeMultipartFormWithBoundary) {
		b, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return fmt.Errorf("read all request body err: %w", err)
		}
		if err := req.Body.Close(); err != nil {
			return fmt.Errorf("request body close err: %w", err)
		}

		req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		req.Header.Set(HTTPHeaderContentMD5, MD5(b))
	}

	stringToSign, err := buildStringToSign(req)
	if err != nil {
		return fmt.Errorf("build string to sign err: %w", err)
	}

	req.Header.Set(HTTPHeaderCASignature, HmacSHA256([]byte(stringToSign), []byte(appKeySecret)))

	return nil
}

func buildStringToSign(req *http.Request) (string, error) {
	var s strings.Builder
	s.WriteString(strings.ToUpper(req.Method) + defaultLF)

	s.WriteString(req.Header.Get(HTTPHeaderAccept) + defaultLF)
	s.WriteString(req.Header.Get(HTTPHeaderContentMD5) + defaultLF)
	s.WriteString(req.Header.Get(HTTPHeaderContentType) + defaultLF)
	s.WriteString(req.Header.Get(HTTPHeaderDate) + defaultLF)

	s.WriteString(buildHeaderStringToSign(req))

	paramStr, err := buildParamStringToSign(req)
	if err != nil {
		return "", fmt.Errorf("build param string to sign err: %w", err)
	}

	s.WriteString(paramStr)

	return s.String(), nil
}

func buildHeaderStringToSign(req *http.Request) string {
	var builder strings.Builder
	signHeaderKeys := make([]string, 0)
	headerKeys := getSortKeys(req.Header, true)

	for _, key := range headerKeys {
		if _, ok := signHeaders[key]; ok {
			signHeaderKeys = append(signHeaderKeys, key)
			builder.WriteString(key + ":" + req.Header.Get(key) + defaultLF)
		}
	}

	req.Header.Set(HTTPHeaderCASignatureHeaders, strings.Join(signHeaderKeys, defaultSep))

	return builder.String()
}

func buildParamStringToSign(req *http.Request) (string, error) {
	clone, err := copyRequest(req)
	if err != nil {
		return "", fmt.Errorf("copy request err: %w", err)
	}

	if err := clone.ParseForm(); err != nil {
		return "", fmt.Errorf("parse form err: %w", err)
	}

	paramKeys := getSortKeys(clone.Form)
	paramList := make([]string, 0)
	for _, key := range paramKeys {
		value := clone.Form.Get(key)
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

	return clone.URL.Path + params, nil
}
