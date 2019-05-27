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

func (s *StdReturn) WithCode(code string) {
	s.Code = code
}

func (s *StdReturn) WithMessage(msg string) {
	s.Message = msg
}
