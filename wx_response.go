package wx

type Response interface {
	Success() bool
}

type CommonResponse struct {
	ErrorCodex int64 `json:"errcode"`
	ErrorMessagex string `json:"errmsg"`
}

func (r CommonResponse) Success() bool {
	return r.ErrorCodex == 0
}