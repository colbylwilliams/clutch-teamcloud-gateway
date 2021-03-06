// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: teamcloud/v1/organization.proto

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

// Validate checks the field values on Organization with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Organization) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Organization with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrganizationMultiError, or
// nil if none found.
func (m *Organization) ValidateAll() error {
	return m.validate(true)
}

func (m *Organization) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Tenant

	// no validation rules for Slug

	// no validation rules for DisplayName

	// no validation rules for SubscriptionId

	// no validation rules for Location

	// no validation rules for Tags

	// no validation rules for ResourceState

	// no validation rules for Id

	if m.ResourceId != nil {
		// no validation rules for ResourceId
	}

	if m.SecretsVaultId != nil {
		// no validation rules for SecretsVaultId
	}

	if m.GalleryId != nil {
		// no validation rules for GalleryId
	}

	if m.RegistryId != nil {
		// no validation rules for RegistryId
	}

	if m.StorageId != nil {
		// no validation rules for StorageId
	}

	if len(errors) > 0 {
		return OrganizationMultiError(errors)
	}

	return nil
}

// OrganizationMultiError is an error wrapping multiple validation errors
// returned by Organization.ValidateAll() if the designated constraints aren't met.
type OrganizationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrganizationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrganizationMultiError) AllErrors() []error { return m }

// OrganizationValidationError is the validation error returned by
// Organization.Validate if the designated constraints aren't met.
type OrganizationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrganizationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrganizationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrganizationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrganizationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrganizationValidationError) ErrorName() string { return "OrganizationValidationError" }

// Error satisfies the builtin error interface
func (e OrganizationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrganization.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrganizationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrganizationValidationError{}
