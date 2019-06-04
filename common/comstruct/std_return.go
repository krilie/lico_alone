package comstruct

//标准返回不引用error类型
//常规返回类型
var (
	StdSuccess = &StdReturn{Code: "success", Message: "success"}
	StdFailure = &StdReturn{Code: "failure", Message: "failure"}
)

// std return type for http
type StdReturn struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (s *StdReturn) WithCode(code string) *StdReturn {
	s.Code = code
	return s
}

func (s *StdReturn) WithMessage(msg string) *StdReturn {
	s.Message = msg
	return s
}
