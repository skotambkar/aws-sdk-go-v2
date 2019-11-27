// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstanceInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInstanceOutput struct {
	_ struct{} `type:"structure"`

	// An array of key-value pairs containing information about the specified instance.
	Instance *Instance `locationName:"instance" type:"structure"`
}

// String returns the string representation
func (s GetInstanceOutput) String() string {
	return awsutil.Prettify(s)
}
