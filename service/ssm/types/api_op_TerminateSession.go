// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type TerminateSessionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the session to terminate.
	//
	// SessionId is a required field
	SessionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s TerminateSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TerminateSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "TerminateSessionInput"}

	if s.SessionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SessionId"))
	}
	if s.SessionId != nil && len(*s.SessionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SessionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type TerminateSessionOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the session that has been terminated.
	SessionId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s TerminateSessionOutput) String() string {
	return awsutil.Prettify(s)
}
