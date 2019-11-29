// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeStackProvisioningParametersInput struct {
	_ struct{} `type:"structure"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeStackProvisioningParametersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeStackProvisioningParametersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeStackProvisioningParametersInput"}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a DescribeStackProvisioningParameters request.
type DescribeStackProvisioningParametersOutput struct {
	_ struct{} `type:"structure"`

	// The AWS OpsWorks Stacks agent installer's URL.
	AgentInstallerUrl *string `type:"string"`

	// An embedded object that contains the provisioning parameters.
	Parameters map[string]string `type:"map"`
}

// String returns the string representation
func (s DescribeStackProvisioningParametersOutput) String() string {
	return awsutil.Prettify(s)
}