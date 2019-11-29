// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type StartJobInput struct {
	_ struct{} `type:"structure"`

	// JobId is a required field
	JobId *string `location:"uri" locationName:"JobId" type:"string" required:"true"`
}

// String returns the string representation
func (s StartJobInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartJobInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartJobInput"}

	if s.JobId == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartJobOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s StartJobOutput) String() string {
	return awsutil.Prettify(s)
}