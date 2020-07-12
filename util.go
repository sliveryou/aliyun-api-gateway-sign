package sign

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"sort"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
)

func getSortKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func CurrentTimeMillis() string {
	t := time.Now().UnixNano() / 1000000
	return strconv.FormatInt(t, 10)
}

func CurrentGMTDate() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

func UUID4() string {
	u := uuid.NewV4()
	return u.String()
}

func HmacSHA256(b []byte, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(b)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Md5(b []byte) string {
	m := md5.New()
	m.Write(b)
	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}
