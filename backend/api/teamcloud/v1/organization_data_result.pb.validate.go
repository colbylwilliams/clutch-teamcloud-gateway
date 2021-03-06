// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/organization_data_result.proto

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

// Validate checks the field values on OrganizationDataResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OrganizationDataResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrganizationDataResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OrganizationDataResultMultiError, or nil if none found.
func (m *OrganizationDataResult) ValidateAll() error {
	return m.validate(true)
}

func (m *OrganizationDataResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OrganizationDataResultValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OrganizationDataResultValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OrganizationDataResultValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Status != nil {
		// no validation rules for Status
	}

	if m.Location != nil {
		// no validation rules for Location
	}

	if len(errors) > 0 {
		return OrganizationDataResultMultiError(errors)
	}

	return nil
}

// OrganizationDataResultMultiError is an error wrapping multiple validation
// errors returned by OrganizationDataResult.ValidateAll() if the designated
// constraints aren't met.
type OrganizationDataResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrganizationDataResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrganizationDataResultMultiError) AllErrors() []error { return m }

// OrganizationDataResultValidationError is the validation error returned by
// OrganizationDataResult.Validate if the designated constraints aren't met.
type OrganizationDataResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationDataResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationDataResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationDataResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationDataResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationDataResultValidationError) ErrorName() string {
	return "OrganizationDataResultValidationError"
}

// Error satisfies the builtin error interface
func (e OrganizationDataResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganizationDataResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationDataResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationDataResultValidationError{}
