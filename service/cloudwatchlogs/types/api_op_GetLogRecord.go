// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLogRecordInput struct {
	_ struct{} `type:"structure"`

	// The pointer corresponding to the log event record you want to retrieve. You
	// get this from the response of a GetQueryResults operation. In that response,
	// the value of the @ptr field for a log event is the value to use as logRecordPointer
	// to retrieve that complete log event record.
	//
	// LogRecordPointer is a required field
	LogRecordPointer *string `locationName:"logRecordPointer" type:"string" required:"true"`
}

// String returns the string representation
func (s GetLogRecordInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLogRecordInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLogRecordInput"}

	if s.LogRecordPointer == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogRecordPointer"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLogRecordOutput struct {
	_ struct{} `type:"structure"`

	// The requested log event, as a JSON string.
	LogRecord map[string]string `locationName:"logRecord" type:"map"`
}

// String returns the string representation
func (s GetLogRecordOutput) String() string {
	return awsutil.Prettify(s)
}
