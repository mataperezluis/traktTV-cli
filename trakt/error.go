//Package trakt ...
package trakt

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/jingweno/go-sawyer"
)

//ResponseErrorType ...
type ResponseErrorType int

const (
	errorClientError             ResponseErrorType = iota // 400-499
	errorBadRequest              ResponseErrorType = iota // 400
	errorUnauthorized            ResponseErrorType = iota // 401
	errorOneTimePasswordRequired ResponseErrorType = iota // 401
	errorForbidden               ResponseErrorType = iota // 403
	errorTooManyRequests         ResponseErrorType = iota // 403
	errorTooManyLoginAttempts    ResponseErrorType = iota // 403
	errorNotFound                ResponseErrorType = iota // 404
	errorNotAcceptable           ResponseErrorType = iota // 406
	errorUnsupportedMediaType    ResponseErrorType = iota // 414
	errorUnprocessableEntity     ResponseErrorType = iota // 422
	errorServerError             ResponseErrorType = iota // 500-599
	errorInternalServerError     ResponseErrorType = iota // 500
	errorNotImplemented          ResponseErrorType = iota // 501
	errorBadGateway              ResponseErrorType = iota // 502
	errorServiceUnavailable      ResponseErrorType = iota // 503
	errorMissingContentType      ResponseErrorType = iota
	errorUnknownError            ResponseErrorType = iota
)

//ErrorObject ...
type ErrorObject struct {
	Resource string `json:"resource,omitempty"`
	Code     string `json:"code,omitempty"`
	Field    string `json:"field,omitempty"`
	Message  string `json:"message,omitempty"`
}

//Error ...
func (e *ErrorObject) Error() string {
	err := fmt.Sprintf("%v error", e.Code)
	if e.Field != "" {
		err = fmt.Sprintf("%v caused by %v field", err, e.Field)
	}
	err = fmt.Sprintf("%v on %v resource", err, e.Resource)
	if e.Message != "" {
		err = fmt.Sprintf("%v: %v", err, e.Message)
	}

	return err
}

//ResponseError ...
type ResponseError struct {
	Response         *http.Response    `json:"-"`
	Type             ResponseErrorType `json:"-"`
	Message          string            `json:"message,omitempty"`
	Err              string            `json:"error,omitempty"`
	Errors           []ErrorObject     `json:"errors,omitempty"`
	DocumentationURL string            `json:"documentation_url,omitempty"`
}

//Error ...
func (e *ResponseError) Error() string {
	return fmt.Sprintf("%v %v: %d - %s",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode, e.errorMessage())
}

func (e *ResponseError) errorMessage() string {
	messages := []string{}

	if e.Message != "" {
		messages = append(messages, e.Message)
	}

	if e.Err != "" {
		m := fmt.Sprintf("Error: %s", e.Err)
		messages = append(messages, m)
	}

	if len(e.Errors) > 0 {
		m := []string{}
		m = append(m, "\nError summary:")
		for _, e := range e.Errors {
			m = append(m, fmt.Sprintf("\t%s", e.Error()))
		}
		messages = append(messages, strings.Join(m, "\n"))
	}

	if e.DocumentationURL != "" {
		messages = append(messages, fmt.Sprintf("// See: %s", e.DocumentationURL))
	}

	return strings.Join(messages, "\n")
}

//NewResponseError ...
func NewResponseError(resp *sawyer.Response) (err *ResponseError) {
	err = &ResponseError{}

	e := resp.Decode(&err)
	if e != nil {
		err.Message = fmt.Sprintf("Problems parsing error message: %s", e)
	}

	err.Response = resp.Response
	err.Type = getResponseErrorType(err)
	return
}

func getResponseErrorType(err *ResponseError) ResponseErrorType {
	code := err.Response.StatusCode
	header := err.Response.Header

	switch {
	case code == http.StatusBadRequest:
		return errorBadRequest

	case code == http.StatusUnauthorized:
		otp := header.Get("X-GitHub-OTP")
		r := regexp.MustCompile(`(?i)required; (\w+)`)
		if r.MatchString(otp) {
			return errorOneTimePasswordRequired
		}

		return errorUnauthorized

	case code == http.StatusForbidden:
		msg := err.Message
		rr := regexp.MustCompile("(?i)rate limit exceeded")
		if rr.MatchString(msg) {
			return errorTooManyRequests
		}
		lr := regexp.MustCompile("(?i)login attempts exceeded")
		if lr.MatchString(msg) {
			return errorTooManyLoginAttempts
		}

		return errorForbidden

	case code == http.StatusNotFound:
		return errorNotFound

	case code == http.StatusNotAcceptable:
		return errorNotAcceptable

	case code == http.StatusUnsupportedMediaType:
		return errorUnsupportedMediaType

	case code == 422:
		return errorUnprocessableEntity

	case code >= 400 && code <= 499:
		return errorClientError

	case code == http.StatusInternalServerError:
		return errorInternalServerError

	case code == http.StatusNotImplemented:
		return errorNotImplemented

	case code == http.StatusBadGateway:
		return errorBadGateway

	case code == http.StatusServiceUnavailable:
		return errorServiceUnavailable

	case code >= 500 && code <= 599:
		return errorServerError
	}

	return errorUnknownError
}
