package matrix

import "fmt"

const (
	_	= iota;
	ErrorNilMatrix;
	ErrorDimensionMismatch;
	ErrorIllegalIndex;
	ExceptionSingular;
	ExceptionNotSPD;
)

type error struct {
	errorCode	int;
}

type Error interface {
	String() string;
	ErrorCode() int;
	OK() bool;
}

func NewError(errorCode int) *error {
	E := new(error);
	E.errorCode = errorCode;
	return E;
}

func (e *error) String() string	{
	if e == nil {
		return "No error"
	}
	switch e.errorCode {
	case ErrorNilMatrix:
		return "Matrix is nil";
	case ErrorDimensionMismatch:
		return "Input dimensions do not match";
	case ErrorIllegalIndex:
		return "Index out of bounds";
	case ExceptionSingular:
		return "Matrix is singular";
	case ExceptionNotSPD:
		return "Matrix is not positive semidefinite";
	}
	return fmt.Sprintf("Error code %d", e.errorCode);
}
func (e *error) ErrorCode() int	{
	if e == nil {
		return 0
	}
	return e.errorCode
}

func (e *error) OK() bool {
	if e == nil {
		return true;
	}
	return e.ErrorCode() == 0;
}
