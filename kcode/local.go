package kcode

import "fmt"

type localCode struct {
	code     int
	message  string
	httpCode HttpStatusCode
	reason   string
	metadata interface{}
}

// Code returns the integer number of current error code.
func (c localCode) Code() int {
	return c.code
}

// Message returns the brief message for current error code.
func (c localCode) Message() string {
	return c.message
}

func (c localCode) HttpCode() HttpStatusCode {
	return c.httpCode
}

func (c localCode) Reason() string {
	return c.reason
}

func (c localCode) Metadata() interface{} {
	return c.metadata
}

// String returns current error code as a string.
func (c localCode) String() string {
	if c.reason != "" {
		return fmt.Sprintf("Code[%d]-HttpCode[%v]: CodeMsg[%s]-HttpMsg[%s] %s %+v", c.code, c.httpCode, c.message, StatusText(c.httpCode), c.reason, c.metadata)
	}
	if c.message != "" {
		return fmt.Sprintf(`Code[%d]-HttpCode[%v]: CodeMsg[%s]-HttpMsg[%s]`, c.code, c.httpCode, c.message, StatusText(c.httpCode))
	}
	return fmt.Sprintf(`Code[%d]-HttpCode[%v]`, c.code, c.httpCode)
}
