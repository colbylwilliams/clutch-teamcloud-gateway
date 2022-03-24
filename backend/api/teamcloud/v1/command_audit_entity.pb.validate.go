// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/command_audit_entity.proto

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

// Validate checks the field values on CommandAuditEntity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommandAuditEntity) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommandAuditEntity with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommandAuditEntityMultiError, or nil if none found.
func (m *CommandAuditEntity) ValidateAll() error {
	return m.validate(true)
}

func (m *CommandAuditEntity) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RuntimeStatus

	if m.PartitionKey != nil {
		// no validation rules for PartitionKey
	}

	if m.RowKey != nil {
		// no validation rules for RowKey
	}

	if m.Timestamp != nil {
		// no validation rules for Timestamp
	}

	if m.ETag != nil {
		// no validation rules for ETag
	}

	if m.CommandId != nil {
		// no validation rules for CommandId
	}

	if m.OrganizationId != nil {
		// no validation rules for OrganizationId
	}

	if m.CommandJson != nil {
		// no validation rules for CommandJson
	}

	if m.ResultJson != nil {
		// no validation rules for ResultJson
	}

	if m.ProjectId != nil {
		// no validation rules for ProjectId
	}

	if m.UserId != nil {
		// no validation rules for UserId
	}

	if m.ParentId != nil {
		// no validation rules for ParentId
	}

	if m.Command != nil {
		// no validation rules for Command
	}

	if m.ComponentTask != nil {
		// no validation rules for ComponentTask
	}

	if m.CustomStatus != nil {
		// no validation rules for CustomStatus
	}

	if m.Errors != nil {
		// no validation rules for Errors
	}

	if m.Created != nil {
		// no validation rules for Created
	}

	if m.Updated != nil {
		// no validation rules for Updated
	}

	if len(errors) > 0 {
		return CommandAuditEntityMultiError(errors)
	}

	return nil
}

// CommandAuditEntityMultiError is an error wrapping multiple validation errors
// returned by CommandAuditEntity.ValidateAll() if the designated constraints
// aren't met.
type CommandAuditEntityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommandAuditEntityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommandAuditEntityMultiError) AllErrors() []error { return m }

// CommandAuditEntityValidationError is the validation error returned by
// CommandAuditEntity.Validate if the designated constraints aren't met.
type CommandAuditEntityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommandAuditEntityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommandAuditEntityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommandAuditEntityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommandAuditEntityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommandAuditEntityValidationError) ErrorName() string {
	return "CommandAuditEntityValidationError"
}

// Error satisfies the builtin error interface
func (e CommandAuditEntityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommandAuditEntity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommandAuditEntityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommandAuditEntityValidationError{}
