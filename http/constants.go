package http

import (
	"fmt"
	"net/http"
	"time"
)

// Defaults
const (
	DefaultTimeout             = 5 * time.Second
	DefaultDialerTimeout       = 5 * time.Second
	DefaultTLSHandshakeTimeout = 5 * time.Second
	DefaultIdleConnTimeout     = 5 * time.Second
	DefaultMaxIdleConnsPerHost = 100
	DefaultDisableKeepAlives   = false
)

var (
	// DefaultClient does not use `http.DefaultClient` because we set timeout, keep alives
	// and max idle connection per host.
	DefaultClient = NewClient(nil)
)

// MIME type is an alias in order to provide an enumeration of possible / well known values.
//
// @see: https://github.com/labstack/echo/blob/master/echo.go#L148
type MIME string

// NewMIME return a new instance of MIME using the given mime
func NewMIME(mime string) MIME {
	return MIME(mime)
}

// String is the fmt.Stringer implementation
func (m MIME) String() string {
	return string(m)
}

// Mime
const (
	charsetUTF8 = "charset=UTF-8"

	// Common
	MIMEApplicationJSON                  MIME = "application/json"
	MIMEApplicationJSONCharsetUTF8       MIME = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationJavaScript            MIME = "application/javascript"
	MIMEApplicationJavaScriptCharsetUTF8 MIME = MIMEApplicationJavaScript + "; " + charsetUTF8
	MIMEApplicationXML                   MIME = "application/xml"
	MIMEApplicationXMLCharsetUTF8        MIME = MIMEApplicationXML + "; " + charsetUTF8
	MIMETextXML                          MIME = "text/xml"
	MIMETextXMLCharsetUTF8               MIME = MIMETextXML + "; " + charsetUTF8
	MIMEApplicationForm                  MIME = "application/x-www-form-urlencoded"
	MIMEApplicationProtobuf              MIME = "application/protobuf"
	MIMEApplicationMsgpack               MIME = "application/msgpack"
	MIMETextHTML                         MIME = "text/html"
	MIMETextHTMLCharsetUTF8              MIME = MIMETextHTML + "; " + charsetUTF8
	MIMETextPlain                        MIME = "text/plain"
	MIMETextPlainCharsetUTF8             MIME = MIMETextPlain + "; " + charsetUTF8
	MIMEMultipartForm                    MIME = "multipart/form-data"
	MIMEOctetStream                      MIME = "application/octet-stream"
)

// Header type is an alias in order to provide an enumeration of possible / well known values.
//
// @see: https://github.com/labstack/echo/blob/master/echo.go#L175
type Header string

// NewHeader return a new instance of Header using the given header
func NewHeader(header string) Header {
	return Header(header)
}

// String is the fmt.Stringer implementation
func (h Header) String() string {
	return string(h)
}

// Headers
const (
	// Common
	HeaderAccept              Header = "Accept"
	HeaderAcceptEncoding      Header = "Accept-Encoding"
	HeaderAllow               Header = "Allow"
	HeaderAuthorization       Header = "Authorization"
	HeaderContentDisposition  Header = "Content-Disposition"
	HeaderContentEncoding     Header = "Content-Encoding"
	HeaderContentLength       Header = "Content-Length"
	HeaderContentType         Header = "Content-Type"
	HeaderCookie              Header = "Cookie"
	HeaderSetCookie           Header = "Set-Cookie"
	HeaderIfModifiedSince     Header = "If-Modified-Since"
	HeaderLastModified        Header = "Last-Modified"
	HeaderLocation            Header = "Location"
	HeaderUpgrade             Header = "Upgrade"
	HeaderVary                Header = "Vary"
	HeaderWWWAuthenticate     Header = "WWW-Authenticate"
	HeaderXForwardedFor       Header = "X-Forwarded-For"
	HeaderXForwardedProto     Header = "X-Forwarded-Proto"
	HeaderXForwardedProtocol  Header = "X-Forwarded-Protocol"
	HeaderXForwardedSsl       Header = "X-Forwarded-Ssl"
	HeaderXUrlScheme          Header = "X-Url-Scheme"
	HeaderXHTTPMethodOverride Header = "X-HTTP-Method-Override"
	HeaderXRealIP             Header = "X-Real-IP"
	HeaderXRequestID          Header = "X-Request-ID"
	HeaderXRequestedWith      Header = "X-Requested-With"
	HeaderServer              Header = "Server"
	HeaderOrigin              Header = "Origin"

	// Access control
	HeaderAccessControlRequestMethod    Header = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   Header = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      Header = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     Header = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     Header = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials Header = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    Header = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           Header = "Access-Control-Max-Age"

	// Security
	HeaderStrictTransportSecurity Header = "Strict-Transport-Security"
	HeaderXContentTypeOptions     Header = "X-Content-Type-Options"
	HeaderXXSSProtection          Header = "X-XSS-Protection"
	HeaderXFrameOptions           Header = "X-Frame-Options"
	HeaderContentSecurityPolicy   Header = "Content-Security-Policy"
	HeaderXCSRFToken              Header = "X-CSRF-Token"
)

// Status type is an alias in order to provide an enumeration of possible / well known values.
// HTTP status codes as registered with IANA.
//
// See: http://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
// See: net/http/status.go
type Status int

// NewMethod return a new instance of Status using the given status code
func NewStatus(statusCode int) Status {
	return Status(statusCode)
}

// String is the fmt.Stringer implementation
func (s Status) String() string {
	statusCode := s.Int()
	return fmt.Sprintf("%d: %s", statusCode, http.StatusText(statusCode))
}

func (s Status) Int() int {
	return int(s)
}

// IsSuccess check if the status code is between 200 and 300
func (s Status) IsSuccess() bool {
	statusCode := s.Int()
	return statusCode >= 200 && statusCode < 300
}

// Status
const (
	StatusContinue           Status = http.StatusContinue           // RFC 7231, 6.2.1
	StatusSwitchingProtocols Status = http.StatusSwitchingProtocols // RFC 7231, 6.2.2
	StatusProcessing         Status = http.StatusProcessing         // RFC 2518, 10.1

	StatusOK                   Status = http.StatusOK                   // RFC 7231, 6.3.1
	StatusCreated              Status = http.StatusCreated              // RFC 7231, 6.3.2
	StatusAccepted             Status = http.StatusAccepted             // RFC 7231, 6.3.3
	StatusNonAuthoritativeInfo Status = http.StatusNonAuthoritativeInfo // RFC 7231, 6.3.4
	StatusNoContent            Status = http.StatusNoContent            // RFC 7231, 6.3.5
	StatusResetContent         Status = http.StatusResetContent         // RFC 7231, 6.3.6
	StatusPartialContent       Status = http.StatusPartialContent       // RFC 7233, 4.1
	StatusMultiStatus          Status = http.StatusMultiStatus          // RFC 4918, 11.1
	StatusAlreadyReported      Status = http.StatusAlreadyReported      // RFC 5842, 7.1
	StatusIMUsed               Status = http.StatusIMUsed               // RFC 3229, 10.4.1

	StatusMultipleChoices   Status = http.StatusMultipleChoices   // RFC 7231, 6.4.1
	StatusMovedPermanently  Status = http.StatusMovedPermanently  // RFC 7231, 6.4.2
	StatusFound             Status = http.StatusFound             // RFC 7231, 6.4.3
	StatusSeeOther          Status = http.StatusSeeOther          // RFC 7231, 6.4.4
	StatusNotModified       Status = http.StatusNotModified       // RFC 7232, 4.1
	StatusUseProxy          Status = http.StatusUseProxy          // RFC 7231, 6.4.5
	_                              = 306                          // RFC 7231, 6.4.6 (Unused)
	StatusTemporaryRedirect Status = http.StatusTemporaryRedirect // RFC 7231, 6.4.7
	StatusPermanentRedirect Status = http.StatusPermanentRedirect // RFC 7538, 3

	StatusBadRequest                   Status = http.StatusBadRequest                   // RFC 7231, 6.5.1
	StatusUnauthorized                 Status = http.StatusUnauthorized                 // RFC 7235, 3.1
	StatusPaymentRequired              Status = http.StatusPaymentRequired              // RFC 7231, 6.5.2
	StatusForbidden                    Status = http.StatusForbidden                    // RFC 7231, 6.5.3
	StatusNotFound                     Status = http.StatusNotFound                     // RFC 7231, 6.5.4
	StatusMethodNotAllowed             Status = http.StatusMethodNotAllowed             // RFC 7231, 6.5.5
	StatusNotAcceptable                Status = http.StatusNotAcceptable                // RFC 7231, 6.5.6
	StatusProxyAuthRequired            Status = http.StatusProxyAuthRequired            // RFC 7235, 3.2
	StatusRequestTimeout               Status = http.StatusRequestTimeout               // RFC 7231, 6.5.7
	StatusConflict                     Status = http.StatusConflict                     // RFC 7231, 6.5.8
	StatusGone                         Status = http.StatusGone                         // RFC 7231, 6.5.9
	StatusLengthRequired               Status = http.StatusLengthRequired               // RFC 7231, 6.5.10
	StatusPreconditionFailed           Status = http.StatusPreconditionFailed           // RFC 7232, 4.2
	StatusRequestEntityTooLarge        Status = http.StatusRequestEntityTooLarge        // RFC 7231, 6.5.11
	StatusRequestURITooLong            Status = http.StatusRequestURITooLong            // RFC 7231, 6.5.12
	StatusUnsupportedMediaType         Status = http.StatusUnsupportedMediaType         // RFC 7231, 6.5.13
	StatusRequestedRangeNotSatisfiable Status = http.StatusRequestedRangeNotSatisfiable // RFC 7233, 4.4
	StatusExpectationFailed            Status = http.StatusExpectationFailed            // RFC 7231, 6.5.14
	StatusTeapot                       Status = http.StatusTeapot                       // RFC 7168, 2.3.3
	StatusUnprocessableEntity          Status = http.StatusUnprocessableEntity          // RFC 4918, 11.2
	StatusLocked                       Status = http.StatusLocked                       // RFC 4918, 11.3
	StatusFailedDependency             Status = http.StatusFailedDependency             // RFC 4918, 11.4
	StatusUpgradeRequired              Status = http.StatusUpgradeRequired              // RFC 7231, 6.5.15
	StatusPreconditionRequired         Status = http.StatusPreconditionRequired         // RFC 6585, 3
	StatusTooManyRequests              Status = http.StatusTooManyRequests              // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  Status = http.StatusRequestHeaderFieldsTooLarge  // RFC 6585, 5
	StatusUnavailableForLegalReasons   Status = http.StatusUnavailableForLegalReasons   // RFC 7725, 3

	StatusInternalServerError           Status = http.StatusInternalServerError           // RFC 7231, 6.6.1
	StatusNotImplemented                Status = http.StatusNotImplemented                // RFC 7231, 6.6.2
	StatusBadGateway                    Status = http.StatusBadGateway                    // RFC 7231, 6.6.3
	StatusServiceUnavailable            Status = http.StatusServiceUnavailable            // RFC 7231, 6.6.4
	StatusGatewayTimeout                Status = http.StatusGatewayTimeout                // RFC 7231, 6.6.5
	StatusHTTPVersionNotSupported       Status = http.StatusHTTPVersionNotSupported       // RFC 7231, 6.6.6
	StatusVariantAlsoNegotiates         Status = http.StatusVariantAlsoNegotiates         // RFC 2295, 8.1
	StatusInsufficientStorage           Status = http.StatusInsufficientStorage           // RFC 4918, 11.5
	StatusLoopDetected                  Status = http.StatusLoopDetected                  // RFC 5842, 7.2
	StatusNotExtended                   Status = http.StatusNotExtended                   // RFC 2774, 7
	StatusNetworkAuthenticationRequired Status = http.StatusNetworkAuthenticationRequired // RFC 6585, 6
)

// Method type is an alias in order to provide an enumeration of possible / well known values.
// Common HTTP methods. Unless otherwise noted, these are defined in RFC 7231 section 4.3.
//
// See: net/http/method.go
type Method string

// NewMethod return a new instance of Method using the given method
func NewMethod(method string) Method {
	return Method(method)
}

// String is the fmt.Stringer implementation
func (m Method) String() string {
	return string(m)
}

// Method
const (
	MethodGet     Method = http.MethodGet
	MethodHead    Method = http.MethodHead
	MethodPost    Method = http.MethodPost
	MethodPut     Method = http.MethodPut
	MethodPatch   Method = http.MethodPatch // RFC 5789
	MethodDelete  Method = http.MethodDelete
	MethodConnect Method = http.MethodConnect
	MethodOptions Method = http.MethodOptions
	MethodTrace   Method = http.MethodTrace
)
