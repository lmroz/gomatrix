package matrix

const (
	_	= iota;
	ErrorNilMatrix;
	ErrorBadInput;
	ErrorIllegalIndex;
)


type error struct {
	errorCode	int;
	errorString	string;
}

type Error interface {
	String() string;
}

/*
type errorMatrix struct {
	*matrix;
	errorCode	int;
	errorString	string;
}

func (A *errorMatrix) ErrorCode() int {
	if A == nil {
		return ErrorNilMatrix
	}
	return A.errorCode;
}

func (A *errorMatrix) ErrorString() string {
	if A == nil {
		return "Matrix is nil"
	}
	return A.errorString;
}
*/

func NewError(errorCode int, errorString string) Error {
	E := new(error);
	E.errorCode = errorCode;
	E.errorString = errorString;
	return E;
}

func (e *error) String() string	{ return e.errorString }

