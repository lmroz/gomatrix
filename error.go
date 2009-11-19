package matrix

const (
	_	= iota;
	ErrorNilMatrix;
	ErrorBadInput;
	ErrorIllegalIndex;
)

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