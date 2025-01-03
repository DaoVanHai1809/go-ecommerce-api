package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
	ErrCodeInvalidToken = 30001
)

// message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrCodeInvalidToken: "token is invalid",
}