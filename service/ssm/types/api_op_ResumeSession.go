// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ResumeSessionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the disconnected session to resume.
	//
	// SessionId is a required field
	SessionId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ResumeSessionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResumeSessionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ResumeSessionInput"}

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

type ResumeSessionOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the session.
	SessionId *string `min:"1" type:"string"`

	// A URL back to SSM Agent on the instance that the Session Manager client uses
	// to send commands and receive output from the instance. Format: wss://ssmmessages.region.amazonaws.com/v1/data-channel/session-id?stream=(input|output).
	//
	// region represents the Region identifier for an AWS Region supported by AWS
	// Systems Manager, such as us-east-2 for the US East (Ohio) Region. For a list
	// of supported region values, see the Region column in the AWS Systems Manager
	// table of regions and endpoints (http://docs.aws.amazon.com/general/latest/gr/rande.html#ssm_region)
	// in the AWS General Reference.
	//
	// session-id represents the ID of a Session Manager session, such as 1a2b3c4dEXAMPLE.
	StreamUrl *string `type:"string"`

	// An encrypted token value containing session and caller information. Used
	// to authenticate the connection to the instance.
	TokenValue *string `type:"string"`
}

// String returns the string representation
func (s ResumeSessionOutput) String() string {
	return awsutil.Prettify(s)
}
