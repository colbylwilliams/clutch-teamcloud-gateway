// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/command_audit_entity_list_data_result.proto

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

// Validate checks the field values on CommandAuditEntityListDataResult with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *CommandAuditEntityListDataResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommandAuditEntityListDataResult with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CommandAuditEntityListDataResultMultiError, or nil if none found.
func (m *CommandAuditEntityListDataResult) ValidateAll() error {
	return m.validate(true)
}

func (m *CommandAuditEntityListDataResult) validate(all bool) error {
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
					errors = append(errors, CommandAuditEntityListDataResultValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CommandAuditEntityListDataResultValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CommandAuditEntityListDataResultValidationError{
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
		return CommandAuditEntityListDataResultMultiError(errors)
	}

	return nil
}

// CommandAuditEntityListDataResultMultiError is an error wrapping multiple
// validation errors returned by
// CommandAuditEntityListDataResult.ValidateAll() if the designated
// constraints aren't met.
type CommandAuditEntityListDataResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommandAuditEntityListDataResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommandAuditEntityListDataResultMultiError) AllErrors() []error { return m }

// CommandAuditEntityListDataResultValidationError is the validation error
// returned by CommandAuditEntityListDataResult.Validate if the designated
// constraints aren't met.
type CommandAuditEntityListDataResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommandAuditEntityListDataResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommandAuditEntityListDataResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommandAuditEntityListDataResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommandAuditEntityListDataResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommandAuditEntityListDataResultValidationError) ErrorName() string {
	return "CommandAuditEntityListDataResultValidationError"
}

// Error satisfies the builtin error interface
func (e CommandAuditEntityListDataResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommandAuditEntityListDataResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommandAuditEntityListDataResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommandAuditEntityListDataResultValidationError{}
