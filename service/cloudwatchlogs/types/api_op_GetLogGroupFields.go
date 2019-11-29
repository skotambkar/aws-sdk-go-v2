// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetLogGroupFieldsInput struct {
	_ struct{} `type:"structure"`

	// The name of the log group to search.
	//
	// LogGroupName is a required field
	LogGroupName *string `locationName:"logGroupName" min:"1" type:"string" required:"true"`

	// The time to set as the center of the query. If you specify time, the 8 minutes
	// before and 8 minutes after this time are searched. If you omit time, the
	// past 15 minutes are queried.
	//
	// The time value is specified as epoch time, the number of seconds since January
	// 1, 1970, 00:00:00 UTC.
	Time *int64 `locationName:"time" type:"long"`
}

// String returns the string representation
func (s GetLogGroupFieldsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLogGroupFieldsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLogGroupFieldsInput"}

	if s.LogGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LogGroupName"))
	}
	if s.LogGroupName != nil && len(*s.LogGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LogGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLogGroupFieldsOutput struct {
	_ struct{} `type:"structure"`

	// The array of fields found in the query. Each object in the array contains
	// the name of the field, along with the percentage of time it appeared in the
	// log events that were queried.
	LogGroupFields []LogGroupField `locationName:"logGroupFields" type:"list"`
}

// String returns the string representation
func (s GetLogGroupFieldsOutput) String() string {
	return awsutil.Prettify(s)
}