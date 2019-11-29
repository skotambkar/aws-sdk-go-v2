// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeProductInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The product identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeProductInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProductInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProductInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeProductOutput struct {
	_ struct{} `type:"structure"`

	// Information about the associated budgets.
	Budgets []BudgetDetail `type:"list"`

	// Summary information about the product view.
	ProductViewSummary *ProductViewSummary `type:"structure"`

	// Information about the provisioning artifacts for the specified product.
	ProvisioningArtifacts []ProvisioningArtifact `type:"list"`
}

// String returns the string representation
func (s DescribeProductOutput) String() string {
	return awsutil.Prettify(s)
}
