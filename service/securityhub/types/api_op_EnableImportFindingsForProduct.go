// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type EnableImportFindingsForProductInput struct {
	_ struct{} `type:"structure"`

	// The ARN of the product to enable the integration for.
	//
	// ProductArn is a required field
	ProductArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableImportFindingsForProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableImportFindingsForProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "EnableImportFindingsForProductInput"}

	if s.ProductArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type EnableImportFindingsForProductOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of your subscription to the product to enable integrations for.
	ProductSubscriptionArn *string `type:"string"`
}

// String returns the string representation
func (s EnableImportFindingsForProductOutput) String() string {
	return awsutil.Prettify(s)
}
