// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTemplatePermissionsInput struct {
	_ struct{} `type:"structure"`

	// AWS account ID that contains the template you are describing.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the template.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTemplatePermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTemplatePermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTemplatePermissionsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateId != nil && len(*s.TemplateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeTemplatePermissionsOutput struct {
	_ struct{} `type:"structure"`

	// A list of resource permissions to be set on the template.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The ARN of the template.
	TemplateArn *string `type:"string"`

	// The ID for the template.
	TemplateId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeTemplatePermissionsOutput) String() string {
	return awsutil.Prettify(s)
}