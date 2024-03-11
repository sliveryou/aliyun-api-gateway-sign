package sign

// HTTP header keys.
const (
	HTTPHeaderAccept      = "Accept"
	HTTPHeaderContentMD5  = "Content-MD5"
	HTTPHeaderContentType = "Content-Type"
	HTTPHeaderDate        = "Date"
	HTTPHeaderUserAgent   = "User-Agent"
)

// HTTP header keys used for Aliyun API gateway signature.
const (
	HTTPHeaderCAPrefix           = "X-Ca-"
	HTTPHeaderCAKey              = "X-Ca-Key"
	HTTPHeaderCANonce            = "X-Ca-Nonce"
	HTTPHeaderCASignature        = "X-Ca-Signature"
	HTTPHeaderCASignatureHeaders = "X-Ca-Signature-Headers"
	HTTPHeaderCASignatureMethod  = "X-Ca-Signature-Method"
	HTTPHeaderCATimestamp        = "X-Ca-Timestamp"
)

// HTTP header content-type values.
const (
	HTTPContentTypeForm                      = "application/x-www-form-urlencoded"
	HTTPContentTypeMultipartForm             = "multipart/form-data"
	HTTPContentTypeMultipartFormWithBoundary = "multipart/form-data; boundary="
	HTTPContentTypeStream                    = "application/octet-stream"
	HTTPContentTypeJson                      = "application/json"
	HTTPContentTypeXml                       = "application/xml"
	HTTPContentTypeText                      = "text/plain"
)

// HTTP method values.
const (
	HTTPMethodGet     = "GET"
	HTTPMethodPost    = "POST"
	HTTPMethodPut     = "PUT"
	HTTPMethodDelete  = "DELETE"
	HTTPMethodPatch   = "PATCH"
	HTTPMethodHead    = "HEAD"
	HTTPMethodOptions = "OPTIONS"
)

// default values.
const (
	defaultUserAgent  = "Go-Aliyun-Sign-Client"
	defaultAccept     = "*/*"
	defaultSignMethod = "HmacSHA256"
	defaultLF         = "\n"
	defaultSep        = ","
)
