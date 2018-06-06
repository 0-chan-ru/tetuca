package common

import (
	"errors"
	"net"
	"strconv"
	"strings"

	"github.com/bakape/thumbnailer"
	"github.com/gorilla/websocket"
)

// Commonly used errors
var (
	ErrNameTooLong         = ErrTooLong("name")
	ErrSubjectTooLong      = ErrTooLong("subject")
	ErrPostPasswordTooLong = ErrTooLong("post password")
	ErrBodyTooLong         = ErrTooLong("post body")
	ErrInvalidCreds        = errors.New("invalid login credentials")
	ErrContainsNull        = errors.New("null byte in non-concatenated message")
	ErrBanned              = errors.New("you are banned from this board")
	ErrInvalidCaptcha      = errors.New("invalid captcha provided")

	// The poster is almost certainly spamming
	ErrSpamDected = errors.New("spam detected")
)

// ErrTooLong is passed, when a field exceeds the maximum string length for
// that specific field
type ErrTooLong string

func (e ErrTooLong) Error() string {
	return string(e) + " too long"
}

// ErrInvalidPostID signifies that the post ID passed by the client is invalid
// in some way. In what way exactly should be evident from the API endpoint.
type ErrInvalidPostID uint64

func (e ErrInvalidPostID) Error() string {
	return "invalid post ID: " + strconv.FormatUint(uint64(e), 10)
}

// Returns, if client-caused error can be safely ignored and not logged
func CanIgnoreClientError(err error) bool {
	switch err {
	case ErrBanned, ErrInvalidCaptcha, ErrSpamDected, websocket.ErrCloseSent:
		return true
	}

	switch err.(type) {
	case thumbnailer.ErrUnsupportedMIME, thumbnailer.ErrInvalidImage:
		return true
	case *net.OpError:
		// Ignore client-side connection loss
		s := err.Error()
		for _, suff := range [...]string{
			"connection reset by peer",
			"broken pipe",
		} {
			if strings.HasSuffix(s, suff) {
				return true
			}
		}
	}

	return false
}
