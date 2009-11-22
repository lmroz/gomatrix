package matrix

const (
	_	= iota;
	ErrorNilMatrix;
	ErrorBadInput;
	ErrorIllegalIndex;
	Exception;
)


type error struct {
	errorCode	int;
	errorString	string;
}

type Error interface {
	String() string;
	ErrorCode() int;
	OK() bool;
}

func NewError(errorCode int, errorString string) *error {
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

func (e *error) OK() bool {
	if e == nil {
		return true;
	}
	return e.ErrorCode() == 0;
}
