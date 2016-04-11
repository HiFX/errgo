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

// internalServer represents an error when something unexpected has happened.
type internalServer struct {
	Err
}

// InternalServerf returns an error which satisfies IsInternalServer().
func InternalServerf(format string, args ...interface{}) error {
	return &internalServer{wrap(nil, http.StatusInternalServerError, format, "", args...)}
}

// NewInternalServer returns an error which wraps err that satisfies
// IsInternalServer().
func NewInternalServer(err error, msg string) error {
	return &internalServer{wrap(err, http.StatusInternalServerError, msg, "")}
}

// IsInternalServer reports whether err was created with InternalServerf() or
// NewInternalServer().
func IsInternalServer(err error) bool {
	err = Cause(err)
	_, ok := err.(*internalServer)
	return ok
}

// notFound represents an error when something has not been found.
type notFound struct {
	Err
}

// NotFoundf returns an error which satisfies IsNotFound().
func NotFoundf(format string, args ...interface{}) error {
	return &notFound{wrap(nil, http.StatusNotFound, format, "", args...)}
}

// NewNotFound returns an error which wraps err that satisfies
// IsNotFound().
func NewNotFound(err error, msg string) error {
	return &notFound{wrap(err, http.StatusNotFound, msg, "")}
}

// IsNotFound reports whether err was created with NotFoundf() or
// NewNotFound().
func IsNotFound(err error) bool {
	err = Cause(err)
	_, ok := err.(*notFound)
	return ok
}

// unauthorized represents an error when an operation is unauthorized.
type unauthorized struct {
	Err
}

// Unauthorizedf returns an error which satisfies IsUnauthorized().
func Unauthorizedf(format string, args ...interface{}) error {
	return &unauthorized{wrap(nil, http.StatusUnauthorized, format, "", args...)}
}

// NewUnauthorized returns an error which wraps err and satisfies
// IsUnauthorized().
func NewUnauthorized(err error, msg string) error {
	return &unauthorized{wrap(err, http.StatusUnauthorized, msg, "")}
}

// IsUnauthorized reports whether err was created with Unauthorizedf() or
// NewUnauthorized().
func IsUnauthorized(err error) bool {
	err = Cause(err)
	_, ok := err.(*unauthorized)
	return ok
}

// notImplemented represents an error when something is not
// implemented.
type notImplemented struct {
	Err
}

// NotImplementedf returns an error which satisfies IsNotImplemented().
func NotImplementedf(format string, args ...interface{}) error {
	return &notImplemented{wrap(nil, http.StatusNotImplemented, format, "", args...)}
}

// NewNotImplemented returns an error which wraps err and satisfies
// IsNotImplemented().
func NewNotImplemented(err error, msg string) error {
	return &notImplemented{wrap(err, http.StatusNotImplemented, msg, "")}
}

// IsNotImplemented reports whether err was created with
// NotImplementedf() or NewNotImplemented().
func IsNotImplemented(err error) bool {
	err = Cause(err)
	_, ok := err.(*notImplemented)
	return ok
}

// badRequest represents an error when a request has bad parameters.
type badRequest struct {
	Err
}

// BadRequestf returns an error which satisfies IsBadRequest().
func BadRequestf(format string, args ...interface{}) error {
	return &badRequest{wrap(nil, http.StatusBadRequest, format, "", args...)}
}

// NewBadRequest returns an error which wraps err that satisfies
// IsBadRequest().
func NewBadRequest(err error, msg string) error {
	return &badRequest{wrap(err, http.StatusBadRequest, msg, "")}
}

// IsBadRequest reports whether err was created with BadRequestf() or
// NewBadRequest().
func IsBadRequest(err error) bool {
	err = Cause(err)
	_, ok := err.(*badRequest)
	return ok
}

// methodNotAllowed represents an error when an HTTP request
// is made with an inappropriate method.
type methodNotAllowed struct {
	Err
}

// MethodNotAllowedf returns an error which satisfies IsMethodNotAllowed().
func MethodNotAllowedf(format string, args ...interface{}) error {
	return &methodNotAllowed{wrap(nil, http.StatusMethodNotAllowed, format, "", args...)}
}

// NewMethodNotAllowed returns an error which wraps err that satisfies
// IsMethodNotAllowed().
func NewMethodNotAllowed(err error, msg string) error {
	return &methodNotAllowed{wrap(err, http.StatusMethodNotAllowed, msg, "")}
}

// IsMethodNotAllowed reports whether err was created with MethodNotAllowedf() or
// NewMethodNotAllowed().
func IsMethodNotAllowed(err error) bool {
	err = Cause(err)
	_, ok := err.(*methodNotAllowed)
	return ok
}
