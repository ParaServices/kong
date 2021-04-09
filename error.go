package kong

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type expectedCodes []int

// ArrayContainsStr ...
func (e expectedCodes) codeMatched(code int) bool {
	for _, value := range e {
		if value == code {
			return true
		}
	}
	return false
}

// KongError ...
type KongError interface {
	Error() string
	Response() *http.Response
	ResponseCode() int
}

type errKongResponse struct {
	expectedResponseCode []int
	response             *http.Response
}

func (e *errKongResponse) Error() string {
	b, err := ioutil.ReadAll(e.response.Body)
	if err != nil {
		return err.Error()
	}
	defer e.response.Body.Close()
	return fmt.Sprintf(
		"Expected response code: %v, actual response code: %v, url: %v, method: %v, resp_body: %v",
		e.expectedResponseCode,
		e.response.StatusCode,
		e.response.Request.URL.String(),
		e.response.Request.Method,
		string(b),
	)
}

// Response ...
func (e *errKongResponse) Response() *http.Response {
	return e.response
}

// ResponseCode ...
func (e *errKongResponse) ResponseCode() int {
	return e.response.StatusCode
}

// NewErrKongResponse ...
func NewErrKongResponse(expectedRespCode []int, resp *http.Response) KongError {
	return &errKongResponse{
		expectedResponseCode: expectedRespCode,
		response:             resp,
	}
}
