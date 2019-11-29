// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package enums

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeResourceNotFoundException   ErrorCode = "ResourceNotFoundException"
	ErrorCodeInvalidRequestException     ErrorCode = "InvalidRequestException"
	ErrorCodeInternalFailureException    ErrorCode = "InternalFailureException"
	ErrorCodeServiceUnavailableException ErrorCode = "ServiceUnavailableException"
	ErrorCodeThrottlingException         ErrorCode = "ThrottlingException"
)

func (enum ErrorCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ErrorCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}