// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/organization_users_service.proto

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

// Validate checks the field values on CreateOrganizationUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateOrganizationUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateOrganizationUserRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// CreateOrganizationUserRequestMultiError, or nil if none found.
func (m *CreateOrganizationUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateOrganizationUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrganizationId

	if all {
		switch v := interface{}(m.GetUserDefinition()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateOrganizationUserRequestValidationError{
					field:  "UserDefinition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateOrganizationUserRequestValidationError{
					field:  "UserDefinition",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserDefinition()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateOrganizationUserRequestValidationError{
				field:  "UserDefinition",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateOrganizationUserRequestMultiError(errors)
	}

	return nil
}

// CreateOrganizationUserRequestMultiError is an error wrapping multiple
// validation errors returned by CreateOrganizationUserRequest.ValidateAll()
// if the designated constraints aren't met.
type CreateOrganizationUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateOrganizationUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateOrganizationUserRequestMultiError) AllErrors() []error { return m }

// CreateOrganizationUserRequestValidationError is the validation error
// returned by CreateOrganizationUserRequest.Validate if the designated
// constraints aren't met.
type CreateOrganizationUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateOrganizationUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateOrganizationUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateOrganizationUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateOrganizationUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateOrganizationUserRequestValidationError) ErrorName() string {
	return "CreateOrganizationUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateOrganizationUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateOrganizationUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateOrganizationUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateOrganizationUserRequestValidationError{}

// Validate checks the field values on DeleteOrganizationUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteOrganizationUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteOrganizationUserRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// DeleteOrganizationUserRequestMultiError, or nil if none found.
func (m *DeleteOrganizationUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteOrganizationUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for OrganizationId

	if len(errors) > 0 {
		return DeleteOrganizationUserRequestMultiError(errors)
	}

	return nil
}

// DeleteOrganizationUserRequestMultiError is an error wrapping multiple
// validation errors returned by DeleteOrganizationUserRequest.ValidateAll()
// if the designated constraints aren't met.
type DeleteOrganizationUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteOrganizationUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteOrganizationUserRequestMultiError) AllErrors() []error { return m }

// DeleteOrganizationUserRequestValidationError is the validation error
// returned by DeleteOrganizationUserRequest.Validate if the designated
// constraints aren't met.
type DeleteOrganizationUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteOrganizationUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteOrganizationUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteOrganizationUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteOrganizationUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteOrganizationUserRequestValidationError) ErrorName() string {
	return "DeleteOrganizationUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteOrganizationUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteOrganizationUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteOrganizationUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteOrganizationUserRequestValidationError{}

// Validate checks the field values on GetOrganizationUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrganizationUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrganizationUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrganizationUserRequestMultiError, or nil if none found.
func (m *GetOrganizationUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrganizationUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for OrganizationId

	if len(errors) > 0 {
		return GetOrganizationUserRequestMultiError(errors)
	}

	return nil
}

// GetOrganizationUserRequestMultiError is an error wrapping multiple
// validation errors returned by GetOrganizationUserRequest.ValidateAll() if
// the designated constraints aren't met.
type GetOrganizationUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrganizationUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrganizationUserRequestMultiError) AllErrors() []error { return m }

// GetOrganizationUserRequestValidationError is the validation error returned
// by GetOrganizationUserRequest.Validate if the designated constraints aren't met.
type GetOrganizationUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrganizationUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrganizationUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrganizationUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrganizationUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrganizationUserRequestValidationError) ErrorName() string {
	return "GetOrganizationUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrganizationUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrganizationUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrganizationUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrganizationUserRequestValidationError{}

// Validate checks the field values on GetOrganizationUserMeRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrganizationUserMeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrganizationUserMeRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrganizationUserMeRequestMultiError, or nil if none found.
func (m *GetOrganizationUserMeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrganizationUserMeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrganizationId

	if len(errors) > 0 {
		return GetOrganizationUserMeRequestMultiError(errors)
	}

	return nil
}

// GetOrganizationUserMeRequestMultiError is an error wrapping multiple
// validation errors returned by GetOrganizationUserMeRequest.ValidateAll() if
// the designated constraints aren't met.
type GetOrganizationUserMeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrganizationUserMeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrganizationUserMeRequestMultiError) AllErrors() []error { return m }

// GetOrganizationUserMeRequestValidationError is the validation error returned
// by GetOrganizationUserMeRequest.Validate if the designated constraints
// aren't met.
type GetOrganizationUserMeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrganizationUserMeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrganizationUserMeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrganizationUserMeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrganizationUserMeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrganizationUserMeRequestValidationError) ErrorName() string {
	return "GetOrganizationUserMeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrganizationUserMeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrganizationUserMeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrganizationUserMeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrganizationUserMeRequestValidationError{}

// Validate checks the field values on GetOrganizationUsersRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetOrganizationUsersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetOrganizationUsersRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetOrganizationUsersRequestMultiError, or nil if none found.
func (m *GetOrganizationUsersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetOrganizationUsersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrganizationId

	if len(errors) > 0 {
		return GetOrganizationUsersRequestMultiError(errors)
	}

	return nil
}

// GetOrganizationUsersRequestMultiError is an error wrapping multiple
// validation errors returned by GetOrganizationUsersRequest.ValidateAll() if
// the designated constraints aren't met.
type GetOrganizationUsersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetOrganizationUsersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetOrganizationUsersRequestMultiError) AllErrors() []error { return m }

// GetOrganizationUsersRequestValidationError is the validation error returned
// by GetOrganizationUsersRequest.Validate if the designated constraints
// aren't met.
type GetOrganizationUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetOrganizationUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetOrganizationUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetOrganizationUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetOrganizationUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetOrganizationUsersRequestValidationError) ErrorName() string {
	return "GetOrganizationUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetOrganizationUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetOrganizationUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetOrganizationUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetOrganizationUsersRequestValidationError{}

// Validate checks the field values on UpdateOrganizationUserRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateOrganizationUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateOrganizationUserRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdateOrganizationUserRequestMultiError, or nil if none found.
func (m *UpdateOrganizationUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateOrganizationUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	// no validation rules for OrganizationId

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateOrganizationUserRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateOrganizationUserRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateOrganizationUserRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateOrganizationUserRequestMultiError(errors)
	}

	return nil
}

// UpdateOrganizationUserRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateOrganizationUserRequest.ValidateAll()
// if the designated constraints aren't met.
type UpdateOrganizationUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateOrganizationUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateOrganizationUserRequestMultiError) AllErrors() []error { return m }

// UpdateOrganizationUserRequestValidationError is the validation error
// returned by UpdateOrganizationUserRequest.Validate if the designated
// constraints aren't met.
type UpdateOrganizationUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateOrganizationUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateOrganizationUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateOrganizationUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateOrganizationUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateOrganizationUserRequestValidationError) ErrorName() string {
	return "UpdateOrganizationUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateOrganizationUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateOrganizationUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateOrganizationUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateOrganizationUserRequestValidationError{}

// Validate checks the field values on UpdateOrganizationUserMeRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateOrganizationUserMeRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateOrganizationUserMeRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// UpdateOrganizationUserMeRequestMultiError, or nil if none found.
func (m *UpdateOrganizationUserMeRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateOrganizationUserMeRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrganizationId

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateOrganizationUserMeRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateOrganizationUserMeRequestValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateOrganizationUserMeRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateOrganizationUserMeRequestMultiError(errors)
	}

	return nil
}

// UpdateOrganizationUserMeRequestMultiError is an error wrapping multiple
// validation errors returned by UpdateOrganizationUserMeRequest.ValidateAll()
// if the designated constraints aren't met.
type UpdateOrganizationUserMeRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateOrganizationUserMeRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateOrganizationUserMeRequestMultiError) AllErrors() []error { return m }

// UpdateOrganizationUserMeRequestValidationError is the validation error
// returned by UpdateOrganizationUserMeRequest.Validate if the designated
// constraints aren't met.
type UpdateOrganizationUserMeRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateOrganizationUserMeRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateOrganizationUserMeRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateOrganizationUserMeRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateOrganizationUserMeRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateOrganizationUserMeRequestValidationError) ErrorName() string {
	return "UpdateOrganizationUserMeRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateOrganizationUserMeRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateOrganizationUserMeRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateOrganizationUserMeRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateOrganizationUserMeRequestValidationError{}
