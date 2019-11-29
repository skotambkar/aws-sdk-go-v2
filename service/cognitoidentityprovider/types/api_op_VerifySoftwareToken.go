// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/enums"
)

type VerifySoftwareTokenInput struct {
	_ struct{} `type:"structure"`

	// The access token.
	AccessToken *string `type:"string" sensitive:"true"`

	// The friendly device name.
	FriendlyDeviceName *string `type:"string"`

	// The session which should be passed both ways in challenge-response calls
	// to the service.
	Session *string `min:"20" type:"string"`

	// The one time password computed using the secret code returned by
	//
	// UserCode is a required field
	UserCode *string `min:"6" type:"string" required:"true"`
}

// String returns the string representation
func (s VerifySoftwareTokenInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *VerifySoftwareTokenInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "VerifySoftwareTokenInput"}
	if s.Session != nil && len(*s.Session) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("Session", 20))
	}

	if s.UserCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserCode"))
	}
	if s.UserCode != nil && len(*s.UserCode) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("UserCode", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type VerifySoftwareTokenOutput struct {
	_ struct{} `type:"structure"`

	// The session which should be passed both ways in challenge-response calls
	// to the service.
	Session *string `min:"20" type:"string"`

	// The status of the verify software token.
	Status enums.VerifySoftwareTokenResponseType `type:"string" enum:"true"`
}

// String returns the string representation
func (s VerifySoftwareTokenOutput) String() string {
	return awsutil.Prettify(s)
}