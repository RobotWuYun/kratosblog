// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/kratosblog/server/v1/article.proto

package v1

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

// Validate checks the field values on CategoryListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CategoryListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryListRequestMultiError, or nil if none found.
func (m *CategoryListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return CategoryListRequestMultiError(errors)
	}

	return nil
}

// CategoryListRequestMultiError is an error wrapping multiple validation
// errors returned by CategoryListRequest.ValidateAll() if the designated
// constraints aren't met.
type CategoryListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryListRequestMultiError) AllErrors() []error { return m }

// CategoryListRequestValidationError is the validation error returned by
// CategoryListRequest.Validate if the designated constraints aren't met.
type CategoryListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryListRequestValidationError) ErrorName() string {
	return "CategoryListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryListRequestValidationError{}

// Validate checks the field values on CategoryListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CategoryListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryListReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CategoryListReplyMultiError, or nil if none found.
func (m *CategoryListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CategoryListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CategoryListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CategoryListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CategoryListReplyMultiError(errors)
	}

	return nil
}

// CategoryListReplyMultiError is an error wrapping multiple validation errors
// returned by CategoryListReply.ValidateAll() if the designated constraints
// aren't met.
type CategoryListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryListReplyMultiError) AllErrors() []error { return m }

// CategoryListReplyValidationError is the validation error returned by
// CategoryListReply.Validate if the designated constraints aren't met.
type CategoryListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryListReplyValidationError) ErrorName() string {
	return "CategoryListReplyValidationError"
}

// Error satisfies the builtin error interface
func (e CategoryListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryListReplyValidationError{}

// Validate checks the field values on TagListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TagListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TagListRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TagListRequestMultiError,
// or nil if none found.
func (m *TagListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TagListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return TagListRequestMultiError(errors)
	}

	return nil
}

// TagListRequestMultiError is an error wrapping multiple validation errors
// returned by TagListRequest.ValidateAll() if the designated constraints
// aren't met.
type TagListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TagListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TagListRequestMultiError) AllErrors() []error { return m }

// TagListRequestValidationError is the validation error returned by
// TagListRequest.Validate if the designated constraints aren't met.
type TagListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagListRequestValidationError) ErrorName() string { return "TagListRequestValidationError" }

// Error satisfies the builtin error interface
func (e TagListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagListRequestValidationError{}

// Validate checks the field values on TagListReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TagListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TagListReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TagListReplyMultiError, or
// nil if none found.
func (m *TagListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *TagListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TagListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TagListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TagListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return TagListReplyMultiError(errors)
	}

	return nil
}

// TagListReplyMultiError is an error wrapping multiple validation errors
// returned by TagListReply.ValidateAll() if the designated constraints aren't met.
type TagListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TagListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TagListReplyMultiError) AllErrors() []error { return m }

// TagListReplyValidationError is the validation error returned by
// TagListReply.Validate if the designated constraints aren't met.
type TagListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagListReplyValidationError) ErrorName() string { return "TagListReplyValidationError" }

// Error satisfies the builtin error interface
func (e TagListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTagListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagListReplyValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Page

	// no validation rules for PageSize

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on ListReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListReplyMultiError, or nil
// if none found.
func (m *ListReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Total

	for idx, item := range m.GetList() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListReplyValidationError{
						field:  fmt.Sprintf("List[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListReplyValidationError{
					field:  fmt.Sprintf("List[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListReplyMultiError(errors)
	}

	return nil
}

// ListReplyMultiError is an error wrapping multiple validation errors returned
// by ListReply.ValidateAll() if the designated constraints aren't met.
type ListReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListReplyMultiError) AllErrors() []error { return m }

// ListReplyValidationError is the validation error returned by
// ListReply.Validate if the designated constraints aren't met.
type ListReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListReplyValidationError) ErrorName() string { return "ListReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListReplyValidationError{}

// Validate checks the field values on DetailRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DetailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DetailRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DetailRequestMultiError, or
// nil if none found.
func (m *DetailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DetailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DetailRequestMultiError(errors)
	}

	return nil
}

// DetailRequestMultiError is an error wrapping multiple validation errors
// returned by DetailRequest.ValidateAll() if the designated constraints
// aren't met.
type DetailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DetailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DetailRequestMultiError) AllErrors() []error { return m }

// DetailRequestValidationError is the validation error returned by
// DetailRequest.Validate if the designated constraints aren't met.
type DetailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetailRequestValidationError) ErrorName() string { return "DetailRequestValidationError" }

// Error satisfies the builtin error interface
func (e DetailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetailRequestValidationError{}

// Validate checks the field values on DetailReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DetailReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DetailReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DetailReplyMultiError, or
// nil if none found.
func (m *DetailReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DetailReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDetail()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DetailReplyValidationError{
					field:  "Detail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DetailReplyValidationError{
					field:  "Detail",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDetail()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DetailReplyValidationError{
				field:  "Detail",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DetailReplyMultiError(errors)
	}

	return nil
}

// DetailReplyMultiError is an error wrapping multiple validation errors
// returned by DetailReply.ValidateAll() if the designated constraints aren't met.
type DetailReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DetailReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DetailReplyMultiError) AllErrors() []error { return m }

// DetailReplyValidationError is the validation error returned by
// DetailReply.Validate if the designated constraints aren't met.
type DetailReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetailReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetailReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetailReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetailReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetailReplyValidationError) ErrorName() string { return "DetailReplyValidationError" }

// Error satisfies the builtin error interface
func (e DetailReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetailReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetailReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetailReplyValidationError{}
