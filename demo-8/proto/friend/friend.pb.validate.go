// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: friend/friend.proto

package friend

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

// Validate checks the field values on FriendBaseInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FriendBaseInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FriendBaseInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FriendBaseInfoMultiError,
// or nil if none found.
func (m *FriendBaseInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *FriendBaseInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetPlayerId() <= 999 {
		err := FriendBaseInfoValidationError{
			field:  "PlayerId",
			reason: "value must be greater than 999",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetName()) > 256 {
		err := FriendBaseInfoValidationError{
			field:  "Name",
			reason: "value length must be at most 256 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_FriendBaseInfo_Name_Pattern.MatchString(m.GetName()) {
		err := FriendBaseInfoValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for IsOnline

	if !strings.HasPrefix(m.GetFrame(), "foo") {
		err := FriendBaseInfoValidationError{
			field:  "Frame",
			reason: "value does not have prefix \"foo\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _FriendBaseInfo_Head_InLookup[m.GetHead()]; !ok {
		err := FriendBaseInfoValidationError{
			field:  "Head",
			reason: "value must be in list [1 2 3 4 5]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetModel(); val < 50 || val > 90 {
		err := FriendBaseInfoValidationError{
			field:  "Model",
			reason: "value must be inside range [50, 90]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetTag()); l < 1 || l > 10 {
		err := FriendBaseInfoValidationError{
			field:  "Tag",
			reason: "value length must be between 1 and 10 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Offline

	// no validation rules for FriendDegree

	if m.GetAddType() != 1 {
		err := FriendBaseInfoValidationError{
			field:  "AddType",
			reason: "value must equal LetterAdd",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for BaseLevel

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = FriendBaseInfoValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetX() != true {
		err := FriendBaseInfoValidationError{
			field:  "X",
			reason: "value must equal true",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	_FriendBaseInfo_Xx_Unique := make(map[float32]struct{}, len(m.GetXx()))

	for idx, item := range m.GetXx() {
		_, _ = idx, item

		if _, exists := _FriendBaseInfo_Xx_Unique[item]; exists {
			err := FriendBaseInfoValidationError{
				field:  fmt.Sprintf("Xx[%v]", idx),
				reason: "repeated value must contain unique items",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {
			_FriendBaseInfo_Xx_Unique[item] = struct{}{}
		}

		// no validation rules for Xx[idx]
	}

	if all {
		switch v := interface{}(m.GetBook()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FriendBaseInfoValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FriendBaseInfoValidationError{
					field:  "Book",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBook()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FriendBaseInfoValidationError{
				field:  "Book",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FriendBaseInfoMultiError(errors)
	}

	return nil
}

func (m *FriendBaseInfo) _validateHostname(host string) error {
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

func (m *FriendBaseInfo) _validateEmail(addr string) error {
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

// FriendBaseInfoMultiError is an error wrapping multiple validation errors
// returned by FriendBaseInfo.ValidateAll() if the designated constraints
// aren't met.
type FriendBaseInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FriendBaseInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FriendBaseInfoMultiError) AllErrors() []error { return m }

// FriendBaseInfoValidationError is the validation error returned by
// FriendBaseInfo.Validate if the designated constraints aren't met.
type FriendBaseInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FriendBaseInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FriendBaseInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FriendBaseInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FriendBaseInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FriendBaseInfoValidationError) ErrorName() string { return "FriendBaseInfoValidationError" }

// Error satisfies the builtin error interface
func (e FriendBaseInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFriendBaseInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FriendBaseInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FriendBaseInfoValidationError{}

var _FriendBaseInfo_Name_Pattern = regexp.MustCompile("^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$")

var _FriendBaseInfo_Head_InLookup = map[uint32]struct{}{
	1: {},
	2: {},
	3: {},
	4: {},
	5: {},
}

// Validate checks the field values on RadarSearchPlayerInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RadarSearchPlayerInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RadarSearchPlayerInfo with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RadarSearchPlayerInfoMultiError, or nil if none found.
func (m *RadarSearchPlayerInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *RadarSearchPlayerInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Distance

	// no validation rules for PlayerId

	// no validation rules for BubbleFrame

	// no validation rules for Head

	// no validation rules for HeadFrame

	// no validation rules for NickName

	if len(errors) > 0 {
		return RadarSearchPlayerInfoMultiError(errors)
	}

	return nil
}

// RadarSearchPlayerInfoMultiError is an error wrapping multiple validation
// errors returned by RadarSearchPlayerInfo.ValidateAll() if the designated
// constraints aren't met.
type RadarSearchPlayerInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RadarSearchPlayerInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RadarSearchPlayerInfoMultiError) AllErrors() []error { return m }

// RadarSearchPlayerInfoValidationError is the validation error returned by
// RadarSearchPlayerInfo.Validate if the designated constraints aren't met.
type RadarSearchPlayerInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RadarSearchPlayerInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RadarSearchPlayerInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RadarSearchPlayerInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RadarSearchPlayerInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RadarSearchPlayerInfoValidationError) ErrorName() string {
	return "RadarSearchPlayerInfoValidationError"
}

// Error satisfies the builtin error interface
func (e RadarSearchPlayerInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRadarSearchPlayerInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RadarSearchPlayerInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RadarSearchPlayerInfoValidationError{}

// Validate checks the field values on SnakeEnumRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SnakeEnumRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SnakeEnumRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SnakeEnumRequestMultiError, or nil if none found.
func (m *SnakeEnumRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SnakeEnumRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for What

	// no validation rules for Who

	// no validation rules for Where

	// no validation rules for Revision

	if all {
		switch v := interface{}(m.GetSub()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SnakeEnumRequestValidationError{
					field:  "Sub",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SnakeEnumRequestValidationError{
					field:  "Sub",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSub()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SnakeEnumRequestValidationError{
				field:  "Sub",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SnakeEnumRequestMultiError(errors)
	}

	return nil
}

// SnakeEnumRequestMultiError is an error wrapping multiple validation errors
// returned by SnakeEnumRequest.ValidateAll() if the designated constraints
// aren't met.
type SnakeEnumRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SnakeEnumRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SnakeEnumRequestMultiError) AllErrors() []error { return m }

// SnakeEnumRequestValidationError is the validation error returned by
// SnakeEnumRequest.Validate if the designated constraints aren't met.
type SnakeEnumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SnakeEnumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SnakeEnumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SnakeEnumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SnakeEnumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SnakeEnumRequestValidationError) ErrorName() string { return "SnakeEnumRequestValidationError" }

// Error satisfies the builtin error interface
func (e SnakeEnumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSnakeEnumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SnakeEnumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SnakeEnumRequestValidationError{}

// Validate checks the field values on Book with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Book) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Book with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in BookMultiError, or nil if none found.
func (m *Book) ValidateAll() error {
	return m.validate(true)
}

func (m *Book) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetCreateTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookValidationError{
					field:  "CreateTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreateTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookValidationError{
				field:  "CreateTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BookMultiError(errors)
	}

	return nil
}

// BookMultiError is an error wrapping multiple validation errors returned by
// Book.ValidateAll() if the designated constraints aren't met.
type BookMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookMultiError) AllErrors() []error { return m }

// BookValidationError is the validation error returned by Book.Validate if the
// designated constraints aren't met.
type BookValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookValidationError) ErrorName() string { return "BookValidationError" }

// Error satisfies the builtin error interface
func (e BookValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBook.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookValidationError{}

// Validate checks the field values on SnakeEnumResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SnakeEnumResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SnakeEnumResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SnakeEnumResponseMultiError, or nil if none found.
func (m *SnakeEnumResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SnakeEnumResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SnakeEnumResponseMultiError(errors)
	}

	return nil
}

// SnakeEnumResponseMultiError is an error wrapping multiple validation errors
// returned by SnakeEnumResponse.ValidateAll() if the designated constraints
// aren't met.
type SnakeEnumResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SnakeEnumResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SnakeEnumResponseMultiError) AllErrors() []error { return m }

// SnakeEnumResponseValidationError is the validation error returned by
// SnakeEnumResponse.Validate if the designated constraints aren't met.
type SnakeEnumResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SnakeEnumResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SnakeEnumResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SnakeEnumResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SnakeEnumResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SnakeEnumResponseValidationError) ErrorName() string {
	return "SnakeEnumResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SnakeEnumResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSnakeEnumResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SnakeEnumResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SnakeEnumResponseValidationError{}

// Validate checks the field values on SnakeEnumRequest_SubMessage with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SnakeEnumRequest_SubMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SnakeEnumRequest_SubMessage with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SnakeEnumRequest_SubMessageMultiError, or nil if none found.
func (m *SnakeEnumRequest_SubMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *SnakeEnumRequest_SubMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SubField

	if len(errors) > 0 {
		return SnakeEnumRequest_SubMessageMultiError(errors)
	}

	return nil
}

// SnakeEnumRequest_SubMessageMultiError is an error wrapping multiple
// validation errors returned by SnakeEnumRequest_SubMessage.ValidateAll() if
// the designated constraints aren't met.
type SnakeEnumRequest_SubMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SnakeEnumRequest_SubMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SnakeEnumRequest_SubMessageMultiError) AllErrors() []error { return m }

// SnakeEnumRequest_SubMessageValidationError is the validation error returned
// by SnakeEnumRequest_SubMessage.Validate if the designated constraints
// aren't met.
type SnakeEnumRequest_SubMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SnakeEnumRequest_SubMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SnakeEnumRequest_SubMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SnakeEnumRequest_SubMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SnakeEnumRequest_SubMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SnakeEnumRequest_SubMessageValidationError) ErrorName() string {
	return "SnakeEnumRequest_SubMessageValidationError"
}

// Error satisfies the builtin error interface
func (e SnakeEnumRequest_SubMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSnakeEnumRequest_SubMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SnakeEnumRequest_SubMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SnakeEnumRequest_SubMessageValidationError{}
