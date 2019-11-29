// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of an UpdateTimeToLive operation.
type UpdateTimeToLiveInput struct {
	_ struct{} `type:"structure"`

	// The name of the table to be configured.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`

	// Represents the settings used to enable or disable Time to Live for the specified
	// table.
	//
	// TimeToLiveSpecification is a required field
	TimeToLiveSpecification *TimeToLiveSpecification `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateTimeToLiveInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTimeToLiveInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTimeToLiveInput"}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}

	if s.TimeToLiveSpecification == nil {
		invalidParams.Add(aws.NewErrParamRequired("TimeToLiveSpecification"))
	}
	if s.TimeToLiveSpecification != nil {
		if err := s.TimeToLiveSpecification.Validate(); err != nil {
			invalidParams.AddNested("TimeToLiveSpecification", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateTimeToLiveOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of an UpdateTimeToLive operation.
	TimeToLiveSpecification *TimeToLiveSpecification `type:"structure"`
}

// String returns the string representation
func (s UpdateTimeToLiveOutput) String() string {
	return awsutil.Prettify(s)
}