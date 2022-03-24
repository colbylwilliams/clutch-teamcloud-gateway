// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/deployment_scopes_authorization_service.proto

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

// Validate checks the field values on InitializeAuthorizationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InitializeAuthorizationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InitializeAuthorizationRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// InitializeAuthorizationRequestMultiError, or nil if none found.
func (m *InitializeAuthorizationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InitializeAuthorizationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrganizationId

	// no validation rules for DeploymentScopeId

	if len(errors) > 0 {
		return InitializeAuthorizationRequestMultiError(errors)
	}

	return nil
}

// InitializeAuthorizationRequestMultiError is an error wrapping multiple
// validation errors returned by InitializeAuthorizationRequest.ValidateAll()
// if the designated constraints aren't met.
type InitializeAuthorizationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InitializeAuthorizationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InitializeAuthorizationRequestMultiError) AllErrors() []error { return m }

// InitializeAuthorizationRequestValidationError is the validation error
// returned by InitializeAuthorizationRequest.Validate if the designated
// constraints aren't met.
type InitializeAuthorizationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InitializeAuthorizationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InitializeAuthorizationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InitializeAuthorizationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InitializeAuthorizationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InitializeAuthorizationRequestValidationError) ErrorName() string {
	return "InitializeAuthorizationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InitializeAuthorizationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInitializeAuthorizationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InitializeAuthorizationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InitializeAuthorizationRequestValidationError{}