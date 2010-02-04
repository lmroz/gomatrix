// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import "fmt"

const (
	NoError	= iota;
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

type Error int

func (e Error) String() string {
	switch e {
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
	return fmt.Sprintf("Unknown error code %d", e);
}

func (e Error) OK() bool	{ return e == NoError }
