// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: hello/hello.proto

package hello

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

// Validate checks the field values on Person with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Person) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Person with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PersonMultiError, or nil if none found.
func (m *Person) ValidateAll() error {
	return m.validate(true)
}

func (m *Person) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetId() <= 999 {
		err := PersonValidationError{
			field:  "Id",
			reason: "value must be greater than 999",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = PersonValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetName()) > 30 {
		err := PersonValidationError{
			field:  "Name",
			reason: "value length must be at most 30 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_Person_Name_Pattern.MatchString(m.GetName()) {
		err := PersonValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"[一-龥]\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetHome() == nil {
		err := PersonValidationError{
			field:  "Home",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetHome()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PersonValidationError{
					field:  "Home",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PersonValidationError{
					field:  "Home",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHome()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PersonValidationError{
				field:  "Home",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return PersonMultiError(errors)
	}

	return nil
}

func (m *Person) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *Person) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// PersonMultiError is an error wrapping multiple validation errors returned by
// Person.ValidateAll() if the designated constraints aren't met.
type PersonMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PersonMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PersonMultiError) AllErrors() []error { return m }

// PersonValidationError is the validation error returned by Person.Validate if
// the designated constraints aren't met.
type PersonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PersonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PersonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PersonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PersonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PersonValidationError) ErrorName() string { return "PersonValidationError" }

// Error satisfies the builtin error interface
func (e PersonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerson.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PersonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PersonValidationError{}

var _Person_Name_Pattern = regexp.MustCompile("[一-龥]")

// Validate checks the field values on Person_Location with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Person_Location) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Person_Location with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Person_LocationMultiError, or nil if none found.
func (m *Person_Location) ValidateAll() error {
	return m.validate(true)
}

func (m *Person_Location) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetLat(); val < -90 || val > 90 {
		err := Person_LocationValidationError{
			field:  "Lat",
			reason: "value must be inside range [-90, 90]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetLng(); val < -180 || val > 180 {
		err := Person_LocationValidationError{
			field:  "Lng",
			reason: "value must be inside range [-180, 180]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return Person_LocationMultiError(errors)
	}

	return nil
}

// Person_LocationMultiError is an error wrapping multiple validation errors
// returned by Person_Location.ValidateAll() if the designated constraints
// aren't met.
type Person_LocationMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Person_LocationMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Person_LocationMultiError) AllErrors() []error { return m }

// Person_LocationValidationError is the validation error returned by
// Person_Location.Validate if the designated constraints aren't met.
type Person_LocationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Person_LocationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Person_LocationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Person_LocationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Person_LocationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Person_LocationValidationError) ErrorName() string { return "Person_LocationValidationError" }

// Error satisfies the builtin error interface
func (e Person_LocationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPerson_Location.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Person_LocationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Person_LocationValidationError{}
