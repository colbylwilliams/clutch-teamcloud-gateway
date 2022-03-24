// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/project_template_definition.proto

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

// Validate checks the field values on ProjectTemplateDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProjectTemplateDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProjectTemplateDefinition with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProjectTemplateDefinitionMultiError, or nil if none found.
func (m *ProjectTemplateDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *ProjectTemplateDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DisplayName

	if all {
		switch v := interface{}(m.GetRepository()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProjectTemplateDefinitionValidationError{
					field:  "Repository",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProjectTemplateDefinitionValidationError{
					field:  "Repository",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRepository()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectTemplateDefinitionValidationError{
				field:  "Repository",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProjectTemplateDefinitionMultiError(errors)
	}

	return nil
}

// ProjectTemplateDefinitionMultiError is an error wrapping multiple validation
// errors returned by ProjectTemplateDefinition.ValidateAll() if the
// designated constraints aren't met.
type ProjectTemplateDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProjectTemplateDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProjectTemplateDefinitionMultiError) AllErrors() []error { return m }

// ProjectTemplateDefinitionValidationError is the validation error returned by
// ProjectTemplateDefinition.Validate if the designated constraints aren't met.
type ProjectTemplateDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectTemplateDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectTemplateDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectTemplateDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectTemplateDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectTemplateDefinitionValidationError) ErrorName() string {
	return "ProjectTemplateDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectTemplateDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectTemplateDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectTemplateDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectTemplateDefinitionValidationError{}
