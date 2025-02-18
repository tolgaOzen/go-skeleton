// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: base/v1/service.proto

package basev1

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

// Validate checks the field values on UserCreateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserCreateRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserCreateRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserCreateRequestMultiError, or nil if none found.
func (m *UserCreateRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserCreateRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetId()) > 64 {
		err := UserCreateRequestValidationError{
			field:  "Id",
			reason: "value length must be at most 64 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_UserCreateRequest_Id_Pattern.MatchString(m.GetId()) {
		err := UserCreateRequestValidationError{
			field:  "Id",
			reason: "value does not match regex pattern \"[a-zA-Z0-9-,]+\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetName()) > 64 {
		err := UserCreateRequestValidationError{
			field:  "Name",
			reason: "value length must be at most 64 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserCreateRequestMultiError(errors)
	}

	return nil
}

// UserCreateRequestMultiError is an error wrapping multiple validation errors
// returned by UserCreateRequest.ValidateAll() if the designated constraints
// aren't met.
type UserCreateRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserCreateRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserCreateRequestMultiError) AllErrors() []error { return m }

// UserCreateRequestValidationError is the validation error returned by
// UserCreateRequest.Validate if the designated constraints aren't met.
type UserCreateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserCreateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserCreateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserCreateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserCreateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserCreateRequestValidationError) ErrorName() string {
	return "UserCreateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UserCreateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserCreateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserCreateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserCreateRequestValidationError{}

var _UserCreateRequest_Id_Pattern = regexp.MustCompile("[a-zA-Z0-9-,]+")

// Validate checks the field values on MessageResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *MessageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MessageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// MessageResponseMultiError, or nil if none found.
func (m *MessageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *MessageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return MessageResponseMultiError(errors)
	}

	return nil
}

// MessageResponseMultiError is an error wrapping multiple validation errors
// returned by MessageResponse.ValidateAll() if the designated constraints
// aren't met.
type MessageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MessageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MessageResponseMultiError) AllErrors() []error { return m }

// MessageResponseValidationError is the validation error returned by
// MessageResponse.Validate if the designated constraints aren't met.
type MessageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MessageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MessageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MessageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MessageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MessageResponseValidationError) ErrorName() string { return "MessageResponseValidationError" }

// Error satisfies the builtin error interface
func (e MessageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMessageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MessageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MessageResponseValidationError{}

// Validate checks the field values on UserListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserListRequestMultiError, or nil if none found.
func (m *UserListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UserListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetSize() <= 0 {
		err := UserListRequestValidationError{
			field:  "Size",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPage() < 0 {
		err := UserListRequestValidationError{
			field:  "Page",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UserListRequestMultiError(errors)
	}

	return nil
}

// UserListRequestMultiError is an error wrapping multiple validation errors
// returned by UserListRequest.ValidateAll() if the designated constraints
// aren't met.
type UserListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserListRequestMultiError) AllErrors() []error { return m }

// UserListRequestValidationError is the validation error returned by
// UserListRequest.Validate if the designated constraints aren't met.
type UserListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListRequestValidationError) ErrorName() string { return "UserListRequestValidationError" }

// Error satisfies the builtin error interface
func (e UserListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListRequestValidationError{}

// Validate checks the field values on UserListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UserListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UserListResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UserListResponseMultiError, or nil if none found.
func (m *UserListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UserListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UserListResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UserListResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserListResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UserListResponseMultiError(errors)
	}

	return nil
}

// UserListResponseMultiError is an error wrapping multiple validation errors
// returned by UserListResponse.ValidateAll() if the designated constraints
// aren't met.
type UserListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserListResponseMultiError) AllErrors() []error { return m }

// UserListResponseValidationError is the validation error returned by
// UserListResponse.Validate if the designated constraints aren't met.
type UserListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserListResponseValidationError) ErrorName() string { return "UserListResponseValidationError" }

// Error satisfies the builtin error interface
func (e UserListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserListResponseValidationError{}
