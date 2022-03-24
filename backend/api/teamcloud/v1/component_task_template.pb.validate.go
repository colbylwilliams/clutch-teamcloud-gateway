// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/component_task_template.proto

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

// Validate checks the field values on ComponentTaskTemplate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ComponentTaskTemplate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ComponentTaskTemplate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ComponentTaskTemplateMultiError, or nil if none found.
func (m *ComponentTaskTemplate) ValidateAll() error {
	return m.validate(true)
}

func (m *ComponentTaskTemplate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if m.Id != nil {
		// no validation rules for Id
	}

	if m.DisplayName != nil {
		// no validation rules for DisplayName
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if m.InputJsonSchema != nil {
		// no validation rules for InputJsonSchema
	}

	if m.TypeName != nil {
		// no validation rules for TypeName
	}

	if len(errors) > 0 {
		return ComponentTaskTemplateMultiError(errors)
	}

	return nil
}

// ComponentTaskTemplateMultiError is an error wrapping multiple validation
// errors returned by ComponentTaskTemplate.ValidateAll() if the designated
// constraints aren't met.
type ComponentTaskTemplateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ComponentTaskTemplateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ComponentTaskTemplateMultiError) AllErrors() []error { return m }

// ComponentTaskTemplateValidationError is the validation error returned by
// ComponentTaskTemplate.Validate if the designated constraints aren't met.
type ComponentTaskTemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComponentTaskTemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComponentTaskTemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComponentTaskTemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComponentTaskTemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComponentTaskTemplateValidationError) ErrorName() string {
	return "ComponentTaskTemplateValidationError"
}

// Error satisfies the builtin error interface
func (e ComponentTaskTemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComponentTaskTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComponentTaskTemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComponentTaskTemplateValidationError{}
