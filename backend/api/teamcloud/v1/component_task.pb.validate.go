// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/component_task.proto

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

// Validate checks the field values on ComponentTask with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ComponentTask) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ComponentTask with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ComponentTaskMultiError, or
// nil if none found.
func (m *ComponentTask) ValidateAll() error {
	return m.validate(true)
}

func (m *ComponentTask) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Organization

	// no validation rules for OrganizationName

	// no validation rules for ComponentId

	// no validation rules for ComponentName

	// no validation rules for ProjectId

	// no validation rules for ProjectName

	// no validation rules for Type

	// no validation rules for Created

	// no validation rules for TaskState

	// no validation rules for Id

	if m.RequestedBy != nil {
		// no validation rules for RequestedBy
	}

	if m.ScheduleId != nil {
		// no validation rules for ScheduleId
	}

	if m.TypeName != nil {
		// no validation rules for TypeName
	}

	if m.Started != nil {
		// no validation rules for Started
	}

	if m.Finished != nil {
		// no validation rules for Finished
	}

	if m.InputJson != nil {
		// no validation rules for InputJson
	}

	if m.Output != nil {
		// no validation rules for Output
	}

	if m.ResourceId != nil {
		// no validation rules for ResourceId
	}

	if m.ExitCode != nil {
		// no validation rules for ExitCode
	}

	if len(errors) > 0 {
		return ComponentTaskMultiError(errors)
	}

	return nil
}

// ComponentTaskMultiError is an error wrapping multiple validation errors
// returned by ComponentTask.ValidateAll() if the designated constraints
// aren't met.
type ComponentTaskMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ComponentTaskMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ComponentTaskMultiError) AllErrors() []error { return m }

// ComponentTaskValidationError is the validation error returned by
// ComponentTask.Validate if the designated constraints aren't met.
type ComponentTaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComponentTaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComponentTaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComponentTaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComponentTaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComponentTaskValidationError) ErrorName() string { return "ComponentTaskValidationError" }

// Error satisfies the builtin error interface
func (e ComponentTaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComponentTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComponentTaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComponentTaskValidationError{}
