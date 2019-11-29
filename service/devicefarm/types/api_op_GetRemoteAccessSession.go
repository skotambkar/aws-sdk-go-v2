// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the request to get information about the specified remote access
// session.
type GetRemoteAccessSessionInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the remote access session about which you
	// want to get session information.
	//
	// Arn is a required field
	Arn *string `locationName:"arn" min:"32" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRemoteAccessSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRemoteAccessSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRemoteAccessSessionInput"}

	if s.Arn == nil {
		invalidParams.Add(aws.NewErrParamRequired("Arn"))
	}
	if s.Arn != nil && len(*s.Arn) < 32 {
		invalidParams.Add(aws.NewErrParamMinLen("Arn", 32))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the response from the server that lists detailed information about
// the remote access session.
type GetRemoteAccessSessionOutput struct {
	_ struct{} `type:"structure"`

	// A container that lists detailed information about the remote access session.
	RemoteAccessSession *RemoteAccessSession `locationName:"remoteAccessSession" type:"structure"`
}

// String returns the string representation
func (s GetRemoteAccessSessionOutput) String() string {
	return awsutil.Prettify(s)
}