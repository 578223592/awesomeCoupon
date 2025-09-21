package response

type Common struct {
	Errno  int    `json:"errno"`
	Data   any    `json:"data"`
	ErrMsg string `json:"errMsg"`
}
