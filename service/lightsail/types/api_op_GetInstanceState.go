// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetInstanceStateInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance to get state information about.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetInstanceStateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetInstanceStateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetInstanceStateInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetInstanceStateOutput struct {
	_ struct{} `type:"structure"`

	// The state of the instance.
	State *InstanceState `locationName:"state" type:"structure"`
}

// String returns the string representation
func (s GetInstanceStateOutput) String() string {
	return awsutil.Prettify(s)
}