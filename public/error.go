package public

type CodeNo struct {
	Code    int
	Message string
}

func (err CodeNo) Error() string {
	return err.Message
}

//DecodeErr 解码error
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *CodeNo:
		return typed.Code, typed.Message //自定义code
	default:
	}

	return InternalServerError.Code, err.Error()
}
