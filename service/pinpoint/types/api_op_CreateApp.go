// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateAppInput struct {
	_ struct{} `type:"structure" payload:"CreateApplicationRequest"`

	// Specifies the display name of an application and the tags to associate with
	// the application.
	//
	// CreateApplicationRequest is a required field
	CreateApplicationRequest *CreateApplicationRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateAppInput"}

	if s.CreateApplicationRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("CreateApplicationRequest"))
	}
	if s.CreateApplicationRequest != nil {
		if err := s.CreateApplicationRequest.Validate(); err != nil {
			invalidParams.AddNested("CreateApplicationRequest", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateAppOutput struct {
	_ struct{} `type:"structure" payload:"ApplicationResponse"`

	// Provides information about an application.
	//
	// ApplicationResponse is a required field
	ApplicationResponse *ApplicationResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateAppOutput) String() string {
	return awsutil.Prettify(s)
}
