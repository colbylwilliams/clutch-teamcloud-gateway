// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: echo/v1/echo.proto

package echov1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on SayHelloRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SayHelloRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayHelloRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SayHelloRequestMultiError, or nil if none found.
func (m *SayHelloRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SayHelloRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) < 1 {
		err := SayHelloRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SayHelloRequestMultiError(errors)
	}

	return nil
}

// SayHelloRequestMultiError is an error wrapping multiple validation errors
// returned by SayHelloRequest.ValidateAll() if the designated constraints
// aren't met.
type SayHelloRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayHelloRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayHelloRequestMultiError) AllErrors() []error { return m }

// SayHelloRequestValidationError is the validation error returned by
// SayHelloRequest.Validate if the designated constraints aren't met.
type SayHelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHelloRequestValidationError) ErrorName() string { return "SayHelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e SayHelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHelloRequestValidationError{}

// Validate checks the field values on SayHelloResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SayHelloResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SayHelloResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SayHelloResponseMultiError, or nil if none found.
func (m *SayHelloResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SayHelloResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return SayHelloResponseMultiError(errors)
	}

	return nil
}

// SayHelloResponseMultiError is an error wrapping multiple validation errors
// returned by SayHelloResponse.ValidateAll() if the designated constraints
// aren't met.
type SayHelloResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SayHelloResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SayHelloResponseMultiError) AllErrors() []error { return m }

// SayHelloResponseValidationError is the validation error returned by
// SayHelloResponse.Validate if the designated constraints aren't met.
type SayHelloResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SayHelloResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SayHelloResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SayHelloResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SayHelloResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SayHelloResponseValidationError) ErrorName() string { return "SayHelloResponseValidationError" }

// Error satisfies the builtin error interface
func (e SayHelloResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSayHelloResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SayHelloResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SayHelloResponseValidationError{}
