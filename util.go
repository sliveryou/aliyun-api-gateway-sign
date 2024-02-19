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

	"github.com/google/uuid"
)

func getSortKeys(m map[string][]string) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
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
