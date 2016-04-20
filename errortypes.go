// Copyright 2014 Canonical Ltd.
// Licensed under the LGPLv3, see LICENCE file for details.

package errgo

import (
	"fmt"
	"net/http"
)

// wrap is a helper to construct an *wrapper.
func wrap(err error, code int, format, suffix string, args ...interface{}) Err {
	newErr := Err{
		message:  fmt.Sprintf(format+suffix, args...),
		previous: err,
	}
	newErr.SetLocation(2)
	newErr.SetCode(code)
	return newErr
}

// InternalServer represents an error when something unexpected has happened.

// InternalServerf returns an error which satisfies IsInternalServer().
func InternalServerf(format string, args ...interface{}) error {
	e := wrap(nil, http.StatusInternalServerError, format, "", args...)
	return &e
}

// NewInternalServer returns an error which wraps err that satisfies
// IsInternalServer().
func NewInternalServer(err error, msg string) error {
	e := wrap(err, http.StatusInternalServerError, msg, "")
	return &e
}

// IsInternalServer reports whether err was created with InternalServerf() or
// NewInternalServer().
func IsInternalServer(err error) bool {
	err = Cause(err)
	e, ok := err.(*Err)
	if ok {
		return e.Code() == http.StatusInternalServerError
	}
	return ok
}

// NotFound represents an error when something has not been found.

// NotFoundf returns an error which satisfies IsNotFound().
func NotFoundf(format string, args ...interface{}) error {
	e := wrap(nil, http.StatusNotFound, format, "", args...)
	return &e
}

// NewNotFound returns an error which wraps err that satisfies
// IsNotFound().
func NewNotFound(err error, msg string) error {
	e := wrap(err, http.StatusNotFound, msg, "")
	return &e
}

// IsNotFound reports whether err was created with NotFoundf() or
// NewNotFound().
func IsNotFound(err error) bool {
	err = Cause(err)
	e, ok := err.(*Err)
	if ok {
		return e.Code() == http.StatusNotFound
	}
	return ok
}

// Unauthorized represents an error when an operation is unauthorized.

// Unauthorizedf returns an error which satisfies IsUnauthorized().
func Unauthorizedf(format string, args ...interface{}) error {
	e := wrap(nil, http.StatusUnauthorized, format, "", args...)
	return &e
}

// NewUnauthorized returns an error which wraps err and satisfies
// IsUnauthorized().
func NewUnauthorized(err error, msg string) error {
	e := wrap(err, http.StatusUnauthorized, msg, "")
	return &e
}

// IsUnauthorized reports whether err was created with Unauthorizedf() or
// NewUnauthorized().
func IsUnauthorized(err error) bool {
	err = Cause(err)
	e, ok := err.(*Err)
	if ok {
		return e.Code() == http.StatusUnauthorized
	}
	return ok
}

// NotImplemented represents an error when something is not
// implemented.

// NotImplementedf returns an error which satisfies IsNotImplemented().
func NotImplementedf(format string, args ...interface{}) error {
	e := wrap(nil, http.StatusNotImplemented, format, "", args...)
	return &e
}

// NewNotImplemented returns an error which wraps err and satisfies
// IsNotImplemented().
func NewNotImplemented(err error, msg string) error {
	e := wrap(err, http.StatusNotImplemented, msg, "")
	return &e
}

// IsNotImplemented reports whether err was created with
// NotImplementedf() or NewNotImplemented().
func IsNotImplemented(err error) bool {
	err = Cause(err)
	e, ok := err.(*Err)
	if ok {
		return e.Code() == http.StatusNotImplemented
	}
	return ok
}

// BadRequest represents an error when a request has bad parameters

// BadRequestf returns an error which satisfies IsBadRequest().
func BadRequestf(format string, args ...interface{}) error {
	e := wrap(nil, http.StatusBadRequest, format, "", args...)
	return &e
}

// NewBadRequest returns an error which wraps err that satisfies
// IsBadRequest().
func NewBadRequest(err error, msg string) error {
	e := wrap(err, http.StatusBadRequest, msg, "")
	return &e
}

// IsBadRequest reports whether err was created with BadRequestf() or
// NewBadRequest().
func IsBadRequest(err error) bool {
	err = Cause(err)
	e, ok := err.(*Err)
	if ok {
		return e.Code() == http.StatusBadRequest
	}
	return ok
}

// MethodNotAllowed represents an error when an HTTP request
// is made with an inappropriate method.

// MethodNotAllowedf returns an error which satisfies IsMethodNotAllowed().
func MethodNotAllowedf(format string, args ...interface{}) error {
	e := wrap(nil, http.StatusMethodNotAllowed, format, "", args...)
	return &e
}

// NewMethodNotAllowed returns an error which wraps err that satisfies
// IsMethodNotAllowed().
func NewMethodNotAllowed(err error, msg string) error {
	e := wrap(err, http.StatusMethodNotAllowed, msg, "")
	return &e
}

// IsMethodNotAllowed reports whether err was created with MethodNotAllowedf() or
// NewMethodNotAllowed().
func IsMethodNotAllowed(err error) bool {
	err = Cause(err)
	e, ok := err.(*Err)
	if ok {
		return e.Code() == http.StatusMethodNotAllowed
	}
	return ok
}
