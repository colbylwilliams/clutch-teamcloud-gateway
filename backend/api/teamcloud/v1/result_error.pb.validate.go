// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/result_error.proto

package teamcloudv1

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

// Validate checks the field values on ResultError with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ResultError) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ResultError with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ResultErrorMultiError, or
// nil if none found.
func (m *ResultError) ValidateAll() error {
	return m.validate(true)
}

func (m *ResultError) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	for idx, item := range m.GetErrors() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ResultErrorValidationError{
						field:  fmt.Sprintf("Errors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ResultErrorValidationError{
						field:  fmt.Sprintf("Errors[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResultErrorValidationError{
					field:  fmt.Sprintf("Errors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Message != nil {
		// no validation rules for Message
	}

	if len(errors) > 0 {
		return ResultErrorMultiError(errors)
	}

	return nil
}

// ResultErrorMultiError is an error wrapping multiple validation errors
// returned by ResultError.ValidateAll() if the designated constraints aren't met.
type ResultErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResultErrorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResultErrorMultiError) AllErrors() []error { return m }

// ResultErrorValidationError is the validation error returned by
// ResultError.Validate if the designated constraints aren't met.
type ResultErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultErrorValidationError) ErrorName() string { return "ResultErrorValidationError" }

// Error satisfies the builtin error interface
func (e ResultErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResultError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultErrorValidationError{}
