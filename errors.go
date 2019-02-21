package kong

import "fmt"

// Error ...
type Error interface {
	Error() string
	ResponseCode() int
}

type errKongResponse struct {
	expectedResponseCode int
	actualResponseCode   int
	rawURL               string
}

func (e *errKongResponse) Error() string {
	return fmt.Sprintf(
		"Expected response code: %v, actual response code: %v, query: %v",
		e.expectedResponseCode,
		e.actualResponseCode,
		e.rawURL,
	)
}

// ResponseCode ...
func (e *errKongResponse) ResponseCode() int {
	return e.actualResponseCode
}

// NewErrKongResponse ...
func NewErrKongResponse(expectedRespCode, actualRespCode int, rawURL string) Error {
	return &errKongResponse{
		expectedResponseCode: expectedRespCode,
		actualResponseCode:   actualRespCode,
		rawURL:               rawURL,
	}
}
