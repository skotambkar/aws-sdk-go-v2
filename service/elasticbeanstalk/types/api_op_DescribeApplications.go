// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to describe one or more applications.
type DescribeApplicationsInput struct {
	_ struct{} `type:"structure"`

	// If specified, AWS Elastic Beanstalk restricts the returned descriptions to
	// only include those with the specified names.
	ApplicationNames []string `type:"list"`
}

// String returns the string representation
func (s DescribeApplicationsInput) String() string {
	return awsutil.Prettify(s)
}

// Result message containing a list of application descriptions.
type DescribeApplicationsOutput struct {
	_ struct{} `type:"structure"`

	// This parameter contains a list of ApplicationDescription.
	Applications []ApplicationDescription `type:"list"`
}

// String returns the string representation
func (s DescribeApplicationsOutput) String() string {
	return awsutil.Prettify(s)
}