// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeTemplateAliasInput struct {
	_ struct{} `type:"structure"`

	// The alias name. $PUBLISHED is not supported for template.
	//
	// AliasName is a required field
	AliasName *string `location:"uri" locationName:"AliasName" min:"1" type:"string" required:"true"`

	// AWS account ID that contains the template alias you are describing.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// An ID for the template.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeTemplateAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeTemplateAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeTemplateAliasInput"}

	if s.AliasName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasName"))
	}
	if s.AliasName != nil && len(*s.AliasName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AliasName", 1))
	}

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

type DescribeTemplateAliasOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// Information about the template alias.
	TemplateAlias *TemplateAlias `type:"structure"`
}

// String returns the string representation
func (s DescribeTemplateAliasOutput) String() string {
	return awsutil.Prettify(s)
}
