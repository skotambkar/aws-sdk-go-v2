// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateServiceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the service that you want to update.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// A complex type that contains the new settings for the service.
	//
	// Service is a required field
	Service *ServiceChange `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateServiceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateServiceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateServiceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if s.Service == nil {
		invalidParams.Add(aws.NewErrParamRequired("Service"))
	}
	if s.Service != nil {
		if err := s.Service.Validate(); err != nil {
			invalidParams.AddNested("Service", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateServiceOutput struct {
	_ struct{} `type:"structure"`

	// A value that you can use to determine whether the request completed successfully.
	// To get the status of the operation, see GetOperation.
	OperationId *string `type:"string"`
}

// String returns the string representation
func (s UpdateServiceOutput) String() string {
	return awsutil.Prettify(s)
}