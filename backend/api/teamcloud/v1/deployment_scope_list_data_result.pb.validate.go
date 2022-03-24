// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/deployment_scope_list_data_result.proto

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

// Validate checks the field values on DeploymentScopeListDataResult with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeploymentScopeListDataResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeploymentScopeListDataResult with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DeploymentScopeListDataResultMultiError, or nil if none found.
func (m *DeploymentScopeListDataResult) ValidateAll() error {
	return m.validate(true)
}

func (m *DeploymentScopeListDataResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DeploymentScopeListDataResultValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DeploymentScopeListDataResultValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DeploymentScopeListDataResultValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
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
		return DeploymentScopeListDataResultMultiError(errors)
	}

	return nil
}

// DeploymentScopeListDataResultMultiError is an error wrapping multiple
// validation errors returned by DeploymentScopeListDataResult.ValidateAll()
// if the designated constraints aren't met.
type DeploymentScopeListDataResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeploymentScopeListDataResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeploymentScopeListDataResultMultiError) AllErrors() []error { return m }

// DeploymentScopeListDataResultValidationError is the validation error
// returned by DeploymentScopeListDataResult.Validate if the designated
// constraints aren't met.
type DeploymentScopeListDataResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeploymentScopeListDataResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeploymentScopeListDataResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeploymentScopeListDataResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeploymentScopeListDataResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeploymentScopeListDataResultValidationError) ErrorName() string {
	return "DeploymentScopeListDataResultValidationError"
}

// Error satisfies the builtin error interface
func (e DeploymentScopeListDataResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeploymentScopeListDataResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeploymentScopeListDataResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeploymentScopeListDataResultValidationError{}
