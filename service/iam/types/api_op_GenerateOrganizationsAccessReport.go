// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GenerateOrganizationsAccessReportInput struct {
	_ struct{} `type:"structure"`

	// The path of the AWS Organizations entity (root, OU, or account). You can
	// build an entity path using the known structure of your organization. For
	// example, assume that your account ID is 123456789012 and its parent OU ID
	// is ou-rge0-awsabcde. The organization root ID is r-f6g7h8i9j0example and
	// your organization ID is o-a1b2c3d4e5. Your entity path is o-a1b2c3d4e5/r-f6g7h8i9j0example/ou-rge0-awsabcde/123456789012.
	//
	// EntityPath is a required field
	EntityPath *string `min:"19" type:"string" required:"true"`

	// The identifier of the AWS Organizations service control policy (SCP). This
	// parameter is optional.
	//
	// This ID is used to generate information about when an account principal that
	// is limited by the SCP attempted to access an AWS service.
	OrganizationsPolicyId *string `type:"string"`
}

// String returns the string representation
func (s GenerateOrganizationsAccessReportInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GenerateOrganizationsAccessReportInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GenerateOrganizationsAccessReportInput"}

	if s.EntityPath == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntityPath"))
	}
	if s.EntityPath != nil && len(*s.EntityPath) < 19 {
		invalidParams.Add(aws.NewErrParamMinLen("EntityPath", 19))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GenerateOrganizationsAccessReportOutput struct {
	_ struct{} `type:"structure"`

	// The job identifier that you can use in the GetOrganizationsAccessReport operation.
	JobId *string `min:"36" type:"string"`
}

// String returns the string representation
func (s GenerateOrganizationsAccessReportOutput) String() string {
	return awsutil.Prettify(s)
}
