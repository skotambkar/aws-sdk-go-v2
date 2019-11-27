// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a GetOnPremisesInstance operation.
type GetOnPremisesInstanceInput struct {
	_ struct{} `type:"structure"`

	// The name of the on-premises instance about which to get information.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetOnPremisesInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetOnPremisesInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetOnPremisesInstanceInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a GetOnPremisesInstance operation.
type GetOnPremisesInstanceOutput struct {
	_ struct{} `type:"structure"`

	// Information about the on-premises instance.
	InstanceInfo *InstanceInfo `locationName:"instanceInfo" type:"structure"`
}

// String returns the string representation
func (s GetOnPremisesInstanceOutput) String() string {
	return awsutil.Prettify(s)
}
