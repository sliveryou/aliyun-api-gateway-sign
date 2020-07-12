package sign

// HTTP header keys.
const (
	HTTPHeaderAccept      = "Accept"
	HTTPHeaderContentMD5  = "Content-MD5"
	HTTPHeaderContentType = "Content-Type"
	HTTPHeaderUserAgent   = "User-Agent"
	HTTPHeaderDate        = "Date"
)

// HTTP header keys used for Aliyun API gateway signature.
const (
	HTTPHeaderCAPrefix           = "X-Ca-"
	HTTPHeaderCASignature        = "X-Ca-Signature"
	HTTPHeaderCATimestamp        = "X-Ca-Timestamp"
	HTTPHeaderCANonce            = "X-Ca-Nonce"
	HTTPHeaderCAKey              = "X-Ca-Key"
	HTTPHeaderCASignatureHeaders = "X-Ca-Signature-Headers"
)

// HTTP header content-type values.
const (
	HTTPContentTypeForm   = "application/x-www-form-urlencoded"
	HTTPContentTypeStream = "application/octet-stream"
	HTTPContentTypeJson   = "application/json"
	HTTPContentTypeXml    = "application/xml"
	HTTPContentTypeText   = "application/text"
)

// HTTP method values.
const (
	HTTPMethodGet = "GET"
	HTTPMethodPost = "POST"
	HTTPMethodPut = "PUT"
	HTTPMethodDelete = "DELETE"
	HTTPMethodPatch = "PATCH"
	HTTPMethodHead = "HEAD"
	HTTPMethodOptions = "OPTIONS"
)

// default values.
const (
	defaultUserAgent = "Go-Aliyun-Sign-Client"
	defaultAccept    = "*/*"
	defaultLF        = "\n"
)

