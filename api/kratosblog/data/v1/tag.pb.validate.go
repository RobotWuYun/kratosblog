// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/kratosblog/data/v1/tag.proto

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

// Validate checks the field values on Tag with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Tag) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Tag with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TagMultiError, or nil if none found.
func (m *Tag) ValidateAll() error {
	return m.validate(true)
}

func (m *Tag) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Color

	// no validation rules for ArticleCount

	if len(errors) > 0 {
		return TagMultiError(errors)
	}

	return nil
}

// TagMultiError is an error wrapping multiple validation errors returned by
// Tag.ValidateAll() if the designated constraints aren't met.
type TagMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TagMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TagMultiError) AllErrors() []error { return m }

// TagValidationError is the validation error returned by Tag.Validate if the
// designated constraints aren't met.
type TagValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TagValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TagValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TagValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TagValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TagValidationError) ErrorName() string { return "TagValidationError" }

// Error satisfies the builtin error interface
func (e TagValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTag.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TagValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TagValidationError{}
