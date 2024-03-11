package sign

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var signHeaders = map[string]struct{}{
	http.CanonicalHeaderKey(HTTPHeaderCAKey):             {},
	http.CanonicalHeaderKey(HTTPHeaderCANonce):           {},
	http.CanonicalHeaderKey(HTTPHeaderCASignatureMethod): {},
	http.CanonicalHeaderKey(HTTPHeaderCATimestamp):       {},
}

func getSortKeys(m map[string][]string, needFormat ...bool) []string {
	nf := false
	if len(needFormat) > 0 {
		nf = needFormat[0]
	}

	keys := make([]string, 0, len(m))
	for key := range m {
		if nf {
			key = http.CanonicalHeaderKey(key)
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

// copyRequest returns the copied request.
func copyRequest(r *http.Request) (*http.Request, error) {
	clone := r.Clone(context.Background())
	if r.Body == nil {
		clone.Body = nil
	} else if r.Body == http.NoBody {
		clone.Body = http.NoBody
	} else if r.GetBody != nil {
		body, err := r.GetBody()
		if err != nil {
			return nil, fmt.Errorf("request get body err: %w", err)
		}

		clone.Body = body
	} else {
		var buf bytes.Buffer
		if _, err := buf.ReadFrom(r.Body); err != nil {
			return nil, fmt.Errorf("read from request body err: %w", err)
		}
		if err := r.Body.Close(); err != nil {
			return nil, fmt.Errorf("request body close err: %w", err)
		}

		r.Body = ioutil.NopCloser(&buf)
		clone.Body = ioutil.NopCloser(bytes.NewBuffer(buf.Bytes()))
	}

	return clone, nil
}

// CurrentTimeMillis returns the millisecond representation of the current time.
func CurrentTimeMillis() string {
	t := time.Now().UnixNano() / 1000000

	return strconv.FormatInt(t, 10)
}

// CurrentGMTDate returns the GMT date representation of the current time.
func CurrentGMTDate() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

// UUID4 returns random generated UUID string.
func UUID4() string {
	u, err := uuid.NewRandom()
	if err != nil {
		return ""
	}

	return u.String()
}

// HmacSHA256 returns the string encrypted with HmacSHA256 method.
func HmacSHA256(b, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(b)

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

// MD5 returns the string hashed with MD5 method.
func MD5(b []byte) string {
	m := md5.New()
	m.Write(b)

	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}
