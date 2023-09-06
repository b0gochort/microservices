package customerrors

type HttpErr struct {
	Code    int
	Message string
}

var (
	ErrNotFound        = HttpErr{Code: 404, Message: "no records!"}
	ErrInternal        = HttpErr{Code: 500, Message: "something went wrong!"}
	ErrValidate        = HttpErr{Code: 422, Message: "failed validate"}
	ErrBindReq         = HttpErr{Code: 400, Message: "failed bind"}
	ErrBalance         = HttpErr{Code: 422, Message: "err with balance"}
	ErrMarshal         = HttpErr{Code: 400, Message: "failed marshal"}
	ErrUnMarshal       = HttpErr{Code: 400, Message: "failed unmarshal"}
	ErrRequest         = HttpErr{Code: 400, Message: "error request to service"}
	ErrGetEnginesByids = HttpErr{Code: 400, Message: "error get engines by car_id"}
)

func (e *HttpErr) Error() string {
	if e == nil {
		return ""
	}

	return e.Message
}
