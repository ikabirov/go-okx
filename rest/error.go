package rest

import (
	"encoding/json"
	"fmt"
)

type OKXError struct {
	StatusCode int
	Code       string
	Message    string
}

type okxResp struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

func NewOKXError(status int, resp string) OKXError {
	var r okxResp
	if err := json.Unmarshal([]byte(resp), &r); err != nil {
		return OKXError{
			StatusCode: status,
			Code:       "-1",
			Message:    resp,
		}
	}

	return OKXError{
		StatusCode: status,
		Code:       r.Code,
		Message:    r.Message,
	}
}

var _ error = (*OKXError)(nil)

func (e OKXError) Error() string {
	return fmt.Sprintf("status: %s,code: %s, message: %s", e.StatusCode, e.Code, e.Message)
}
