// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "fmt"

const (
	_	= iota;
	//The matrix returned was nil.
	ErrorNilMatrix;
	//The dimensions of the inputs do not make sense for this operation.
	ErrorDimensionMismatch;
	//The indices provided are out of bounds.
	ErrorIllegalIndex;
	//The matrix provided has a singularity.
	ExceptionSingular;
	//The matrix provided is not positive semi-definite.
	ExceptionNotSPD;
)

type error struct {
	errorCode int;
}

/*
Error types from matrix operations satisfy this interface.
*/
type Error interface {
	//An english string describing the error.
	String() string;
	//The code for the error.
	ErrorCode() int;
	//If OK()==true, there is no error.
	OK() bool;
}

/*
Create a new error with the provided code.
*/
func NewError(errorCode int) *error {
	E := new(error);
	E.errorCode = errorCode;
	return E;
}

func (e *error) String() string {
	if e == nil {
		return "No error"
	}
	switch e.errorCode {
	case ErrorNilMatrix:
		return "Matrix is nil"
	case ErrorDimensionMismatch:
		return "Input dimensions do not match"
	case ErrorIllegalIndex:
		return "Index out of bounds"
	case ExceptionSingular:
		return "Matrix is singular"
	case ExceptionNotSPD:
		return "Matrix is not positive semidefinite"
	}
	return fmt.Sprintf("Error code %d", e.errorCode);
}
func (e *error) ErrorCode() int {
	if e == nil {
		return 0
	}
	return e.errorCode;
}

func (e *error) OK() bool {
	if e == nil {
		return true
	}
	return e.ErrorCode() == 0;
}
