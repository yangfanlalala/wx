package wx

import "errors"

type Response interface {
	Success() bool
	Error() error
}

type CommonResponse struct {
	ErrorCode int64 `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

func (r CommonResponse) Success() bool {
	return r.ErrorCode == 0
}

func (r CommonResponse) Error() error {
	if r.ErrorCode != 0 {
		return errors.New(r.ErrorMessage)
	}
	return nil
}