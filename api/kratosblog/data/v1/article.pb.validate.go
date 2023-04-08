// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/kratosblog/data/v1/article.proto

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

// Validate checks the field values on Article with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Article) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Article with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ArticleMultiError, or nil if none found.
func (m *Article) ValidateAll() error {
	return m.validate(true)
}

func (m *Article) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Title

	if all {
		switch v := interface{}(m.GetCover()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ArticleValidationError{
					field:  "Cover",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ArticleValidationError{
					field:  "Cover",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCover()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ArticleValidationError{
				field:  "Cover",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Content

	for idx, item := range m.GetTags() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ArticleValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ArticleValidationError{
						field:  fmt.Sprintf("Tags[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ArticleValidationError{
					field:  fmt.Sprintf("Tags[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ArticleMultiError(errors)
	}

	return nil
}

// ArticleMultiError is an error wrapping multiple validation errors returned
// by Article.ValidateAll() if the designated constraints aren't met.
type ArticleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ArticleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ArticleMultiError) AllErrors() []error { return m }

// ArticleValidationError is the validation error returned by Article.Validate
// if the designated constraints aren't met.
type ArticleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ArticleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ArticleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ArticleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ArticleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ArticleValidationError) ErrorName() string { return "ArticleValidationError" }

// Error satisfies the builtin error interface
func (e ArticleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sArticle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ArticleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ArticleValidationError{}
