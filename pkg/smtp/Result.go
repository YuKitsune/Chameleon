package smtp

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// Result represents a response to an SMTP client after receiving DATA.
// The String method should return an SMTP message ready to send back to the
// client, for example `250 OK: Message received`.
type Result interface {
	fmt.Stringer
	// Code should return the SMTP code associated with this response, ie. `250`
	Code() int
}

// Internal implementation of BackendResult for use by backend implementations.
type result struct {
	// we're going to use a bytes.Buffer for building a string
	bytes.Buffer
}

func (r *result) String() string {
	return r.Buffer.String()
}

// Code parses the SMTP code from the first 3 characters of the SMTP message.
// Returns 554 if code cannot be parsed.
func (r *result) Code() int {
	trimmed := strings.TrimSpace(r.String())
	if len(trimmed) < 3 {
		return 554
	}
	code, err := strconv.Atoi(trimmed[:3])
	if err != nil {
		return 554
	}
	return code
}

func NewResult(r ...interface{}) Result {
	buf := new(result)
	for _, item := range r {
		switch v := item.(type) {
		case error:
			_, _ = buf.WriteString(v.Error())
		case fmt.Stringer:
			_, _ = buf.WriteString(v.String())
		case string:
			_, _ = buf.WriteString(v)
		}
	}
	return buf
}
