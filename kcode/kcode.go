package kcode

type HttpStatusCode int

// Code is universal error code interface definition.
type Code interface {
    // Code returns the integer number of current error code.
    Code() int

    // Message returns the brief message for current error code. 业务错误的简短信息
    Message() string

    // HttpCode 返回错误绑定的Http状态码
    HttpCode() HttpStatusCode

    // Reason 返回给用户的可读信息，可以作为用户的提示信息用。可以为空
    Reason() string

    // Metadata 错误元信息，为错误添加附加可扩展的信息
    Metadata() interface{}
}

var (
    CodeNil                      = New(-1, HttpStatusNotSpecified, "", "", nil)                                   // No Error code specified
    CodeOK                       = New(0, HttpStatusOK, "OK", "", nil)                                            // It is OK.
    CodeInternalError            = New(50, HttpStatusInternalServerError, "Internal Error", "", nil)              // An error occurred internally.
    CodeValidationFailed         = New(51, HttpStatusBadRequest, "Validation Failed", "", nil)                    // Data validation failed.
    CodeDbOperationError         = New(52, HttpStatusInternalServerError, "Database Operation Error", "", nil)    // Database operation error.
    CodeInvalidParameter         = New(53, HttpStatusBadRequest, "Invalid Parameter", "", nil)                    // The given parameter for current operation is invalid.
    CodeMissingParameter         = New(54, HttpStatusBadRequest, "Missing Parameter", "", nil)                    // Parameter for current operation is missing.
    CodeInvalidOperation         = New(55, HttpStatusMethodNotAllowed, "Invalid Operation", "", nil)              // The function cannot be used like this.
    CodeInvalidConfiguration     = New(56, HttpStatusInternalServerError, "Invalid Configuration", "", nil)       // The configuration is invalid for current operation.
    CodeMissingConfiguration     = New(57, HttpStatusInternalServerError, "Missing Configuration", "", nil)       // The configuration is missing for current operation.
    CodeNotImplemented           = New(58, HttpStatusNotImplemented, "Not Implemented", "", nil)                  // The operation is not implemented yet.
    CodeNotSupported             = New(59, HttpStatusMethodNotAllowed, "Not Supported", "", nil)                  // The operation is not supported yet.
    CodeOperationFailed          = New(60, HttpStatusInternalServerError, "Operation Failed", "", nil)            // I tried, but I cannot give you what you want.
    CodeNotAuthorized            = New(61, HttpStatusUnauthorized, "Not Authorized", "", nil)                     // Not Authorized.
    CodeSecurityReason           = New(62, HttpStatusForbidden, "Security Reason", "", nil)                       // Security Reason.
    CodeServerBusy               = New(63, HttpStatusTooManyRequests, "Server Is Busy", "", nil)                  // Server is busy, please try again later.
    CodeUnknown                  = New(64, HttpStatusInternalServerError, "Unknown Error", "", nil)               // Unknown error.
    CodeNotFound                 = New(65, HttpStatusNotFound, "Not Found", "", nil)                              // Resource does not exist.
    CodeInvalidRequest           = New(66, HttpStatusBadRequest, "Invalid Request", "", nil)                      // Invalid request.
    CodeBusinessValidationFailed = New(300, HttpStatusInternalServerError, "Business Validation Failed", "", nil) // Business validation failed.
)

func New(code int, httpCode HttpStatusCode, message string, reason string, metadata interface{}) Code {
    return localCode{
        code:     code,
        message:  message,
        httpCode: httpCode,
        reason:   reason,
        metadata: metadata,
    }
}

// WithCode creates and returns a new error code based on given Code.
// reason和metadata字段开放给用户填写
func WithCode(code Code, reason string, metadata interface{}) Code {
    return localCode{
        code:     code.Code(),
        message:  code.Message(),
        httpCode: code.HttpCode(),
        reason:   reason,
        metadata: metadata,
    }
}
