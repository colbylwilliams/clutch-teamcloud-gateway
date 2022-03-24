// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/component_template.proto

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

// Validate checks the field values on ComponentTemplate with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ComponentTemplate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ComponentTemplate with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ComponentTemplateMultiError, or nil if none found.
func (m *ComponentTemplate) ValidateAll() error {
	return m.validate(true)
}

func (m *ComponentTemplate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Organization

	// no validation rules for OrganizationName

	// no validation rules for ParentId

	if all {
		switch v := interface{}(m.GetRepository()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ComponentTemplateValidationError{
					field:  "Repository",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ComponentTemplateValidationError{
					field:  "Repository",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRepository()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ComponentTemplateValidationError{
				field:  "Repository",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPermissions()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ComponentTemplateValidationError{
					field:  "Permissions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ComponentTemplateValidationError{
					field:  "Permissions",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPermissions()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ComponentTemplateValidationError{
				field:  "Permissions",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetTasks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ComponentTemplateValidationError{
						field:  fmt.Sprintf("Tasks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ComponentTemplateValidationError{
						field:  fmt.Sprintf("Tasks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ComponentTemplateValidationError{
					field:  fmt.Sprintf("Tasks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetTaskRunner()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ComponentTemplateValidationError{
					field:  "TaskRunner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ComponentTemplateValidationError{
					field:  "TaskRunner",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTaskRunner()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ComponentTemplateValidationError{
				field:  "TaskRunner",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Type

	// no validation rules for Id

	if m.DisplayName != nil {
		// no validation rules for DisplayName
	}

	if m.Description != nil {
		// no validation rules for Description
	}

	if m.InputJsonSchema != nil {
		// no validation rules for InputJsonSchema
	}

	if m.Folder != nil {
		// no validation rules for Folder
	}

	if m.Configuration != nil {
		// no validation rules for Configuration
	}

	if len(errors) > 0 {
		return ComponentTemplateMultiError(errors)
	}

	return nil
}

// ComponentTemplateMultiError is an error wrapping multiple validation errors
// returned by ComponentTemplate.ValidateAll() if the designated constraints
// aren't met.
type ComponentTemplateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ComponentTemplateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ComponentTemplateMultiError) AllErrors() []error { return m }

// ComponentTemplateValidationError is the validation error returned by
// ComponentTemplate.Validate if the designated constraints aren't met.
type ComponentTemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ComponentTemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ComponentTemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ComponentTemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ComponentTemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ComponentTemplateValidationError) ErrorName() string {
	return "ComponentTemplateValidationError"
}

// Error satisfies the builtin error interface
func (e ComponentTemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sComponentTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ComponentTemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ComponentTemplateValidationError{}