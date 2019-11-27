// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input argument to the TerminationProtection operation.
type SetTerminationProtectionInput struct {
	_ struct{} `type:"structure"`

	// A list of strings that uniquely identify the clusters to protect. This identifier
	// is returned by RunJobFlow and can also be obtained from DescribeJobFlows .
	//
	// JobFlowIds is a required field
	JobFlowIds []string `type:"list" required:"true"`

	// A Boolean that indicates whether to protect the cluster and prevent the Amazon
	// EC2 instances in the cluster from shutting down due to API calls, user intervention,
	// or job-flow error.
	//
	// TerminationProtected is a required field
	TerminationProtected *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s SetTerminationProtectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetTerminationProtectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetTerminationProtectionInput"}

	if s.JobFlowIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("JobFlowIds"))
	}

	if s.TerminationProtected == nil {
		invalidParams.Add(aws.NewErrParamRequired("TerminationProtected"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetTerminationProtectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetTerminationProtectionOutput) String() string {
	return awsutil.Prettify(s)
}
