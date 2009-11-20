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
	ErrorCode() int;
}

func NewError(errorCode int, errorString string) Error {
	E := new(error);
	E.errorCode = errorCode;
	E.errorString = errorString;
	return E;
}

func (e *error) String() string	{
	if e == nil {
		return "no error"
	}
	return e.errorString
}
func (e *error) ErrorCode() int	{
	if e == nil {
		return 0
	}
	return e.errorCode
}

