package response

const (
	ErrCodeSuccess      = 20001
	ErrCodeParamInvalid = 20003
)

// message
var msg = map[int]string {
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
}