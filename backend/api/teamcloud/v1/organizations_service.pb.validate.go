// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/organizations_service.proto

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

// Validate checks the field values on CreateOrganizationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrganizationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrganizationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateOrganizationRequestMultiError, or nil if none found.
func (m *CreateOrganizationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrganizationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetOrganizationDefinition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateOrganizationRequestValidationError{
					field:  "OrganizationDefinition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateOrganizationRequestValidationError{
					field:  "OrganizationDefinition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrganizationDefinition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOrganizationRequestValidationError{
				field:  "OrganizationDefinition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateOrganizationRequestMultiError(errors)
	}

	return nil
}

// CreateOrganizationRequestMultiError is an error wrapping multiple validation
// errors returned by CreateOrganizationRequest.ValidateAll() if the
// designated constraints aren't met.
type CreateOrganizationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrganizationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrganizationRequestMultiError) AllErrors() []error { return m }

// CreateOrganizationRequestValidationError is the validation error returned by
// CreateOrganizationRequest.Validate if the designated constraints aren't met.
type CreateOrganizationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrganizationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrganizationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrganizationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrganizationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrganizationRequestValidationError) ErrorName() string {
	return "CreateOrganizationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrganizationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrganizationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrganizationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrganizationRequestValidationError{}

// Validate checks the field values on DeleteOrganizationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteOrganizationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteOrganizationRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteOrganizationRequestMultiError, or nil if none found.
func (m *DeleteOrganizationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteOrganizationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrganizationId

	if len(errors) > 0 {
		return DeleteOrganizationRequestMultiError(errors)
	}

	return nil
}

// DeleteOrganizationRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteOrganizationRequest.ValidateAll() if the
// designated constraints aren't met.
type DeleteOrganizationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteOrganizationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteOrganizationRequestMultiError) AllErrors() []error { return m }

// DeleteOrganizationRequestValidationError is the validation error returned by
// DeleteOrganizationRequest.Validate if the designated constraints aren't met.
type DeleteOrganizationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteOrganizationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteOrganizationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteOrganizationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteOrganizationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteOrganizationRequestValidationError) ErrorName() string {
	return "DeleteOrganizationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteOrganizationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteOrganizationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteOrganizationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteOrganizationRequestValidationError{}

// Validate checks the field values on GetOrganizationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrganizationRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrganizationRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrganizationRequestMultiError, or nil if none found.
func (m *GetOrganizationRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrganizationRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrganizationId

	if len(errors) > 0 {
		return GetOrganizationRequestMultiError(errors)
	}

	return nil
}

// GetOrganizationRequestMultiError is an error wrapping multiple validation
// errors returned by GetOrganizationRequest.ValidateAll() if the designated
// constraints aren't met.
type GetOrganizationRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrganizationRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrganizationRequestMultiError) AllErrors() []error { return m }

// GetOrganizationRequestValidationError is the validation error returned by
// GetOrganizationRequest.Validate if the designated constraints aren't met.
type GetOrganizationRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrganizationRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrganizationRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrganizationRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrganizationRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrganizationRequestValidationError) ErrorName() string {
	return "GetOrganizationRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrganizationRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrganizationRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrganizationRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrganizationRequestValidationError{}