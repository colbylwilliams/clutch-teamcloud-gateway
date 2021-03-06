// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/project_identity_definition.proto

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

// Validate checks the field values on ProjectIdentityDefinition with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProjectIdentityDefinition) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProjectIdentityDefinition with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProjectIdentityDefinitionMultiError, or nil if none found.
func (m *ProjectIdentityDefinition) ValidateAll() error {
	return m.validate(true)
}

func (m *ProjectIdentityDefinition) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for DisplayName

	// no validation rules for DeploymentScopeId

	if len(errors) > 0 {
		return ProjectIdentityDefinitionMultiError(errors)
	}

	return nil
}

// ProjectIdentityDefinitionMultiError is an error wrapping multiple validation
// errors returned by ProjectIdentityDefinition.ValidateAll() if the
// designated constraints aren't met.
type ProjectIdentityDefinitionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProjectIdentityDefinitionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProjectIdentityDefinitionMultiError) AllErrors() []error { return m }

// ProjectIdentityDefinitionValidationError is the validation error returned by
// ProjectIdentityDefinition.Validate if the designated constraints aren't met.
type ProjectIdentityDefinitionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectIdentityDefinitionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectIdentityDefinitionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectIdentityDefinitionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectIdentityDefinitionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectIdentityDefinitionValidationError) ErrorName() string {
	return "ProjectIdentityDefinitionValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectIdentityDefinitionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectIdentityDefinition.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectIdentityDefinitionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectIdentityDefinitionValidationError{}
