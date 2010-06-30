// Copyright 2009 The GoMatrix Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package matrix

import (
	"fmt"
	"os"
)

const (
	noError = iota
	//The matrix returned was nil.
	errorNilMatrix
	//The dimensions of the inputs do not make sense for this operation.
	errorDimensionMismatch
	//The indices provided are out of bounds.
	errorIllegalIndex
	//The matrix provided has a singularity.
	exceptionSingular
	//The matrix provided is not positive semi-definite.
	exceptionNotSPD
)

type Error interface {
	os.Error
	OK() bool
}

type error int

func (e error) String() string {
	switch e {
	case errorNilMatrix:
		return "Matrix is nil"
	case errorDimensionMismatch:
		return "Input dimensions do not match"
	case errorIllegalIndex:
		return "Index out of bounds"
	case exceptionSingular:
		return "Matrix is singular"
	case exceptionNotSPD:
		return "Matrix is not positive semidefinite"
	}
	return fmt.Sprintf("Unknown error code %d", e)
}

func (e error) OK() bool { return e == noError }

var (
	NoError Error = nil
	//The matrix returned was nil.
	ErrorNilMatrix Error = error(errorNilMatrix)
	//The dimensions of the inputs do not make sense for this operation.
	ErrorDimensionMismatch Error = error(errorDimensionMismatch)
	//The indices provided are out of bounds.
	ErrorIllegalIndex Error = error(errorIllegalIndex)
	//The matrix provided has a singularity.
	ExceptionSingular Error = error(exceptionSingular)
	//The matrix provided is not positive semi-definite.
	ExceptionNotSPD Error = error(exceptionNotSPD)
)
