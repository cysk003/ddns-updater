package errors

import "errors"

var (
	ErrCredentialsNotSet       = errors.New("credentials are not set")
	ErrEmptyAPIKey             = errors.New("empty API key")
	ErrEmptyAppKey             = errors.New("empty app key")
	ErrEmptyCustomerNumber     = errors.New("empty customer number")
	ErrEmptyConsumerKey        = errors.New("empty consumer key")
	ErrEmptyEmail              = errors.New("empty email")
	ErrEmptyKey                = errors.New("empty key")
	ErrEmptyName               = errors.New("empty name")
	ErrEmptyPassword           = errors.New("empty password")
	ErrEmptyAPISecret          = errors.New("empty API secret")
	ErrEmptySecret             = errors.New("empty secret")
	ErrEmptyToken              = errors.New("empty token")
	ErrEmptyAccessKeyID        = errors.New("empty access key id")
	ErrEmptyAccessKeySecret    = errors.New("empty key secret")
	ErrEmptyTTL                = errors.New("TTL is not set")
	ErrEmptyUsername           = errors.New("empty username")
	ErrEmptyZoneIdentifier     = errors.New("empty zone identifier")
	ErrEmptyHost               = errors.New("host cannot be empty")
	ErrGCPProjectNotSet        = errors.New("GCP project is not set")
	ErrHostOnlyAt              = errors.New(`host can only be "@"`)
	ErrHostOnlySubdomain       = errors.New("host can only be a subdomain")
	ErrHostWildcard            = errors.New(`host cannot be a "*"`)
	ErrIPv6NotSupported        = errors.New("IPv6 is not supported by this provider")
	ErrMalformedEmail          = errors.New("malformed email address")
	ErrMalformedKey            = errors.New("malformed key")
	ErrMalformedPassword       = errors.New("malformed password")
	ErrMalformedToken          = errors.New("malformed token")
	ErrMalformedUsername       = errors.New("malformed username")
	ErrMalformedUserServiceKey = errors.New("malformed user service key")
)
