// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreatePortfolioShareInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The AWS account ID. For example, 123456789012.
	AccountId *string `type:"string"`

	// The organization node to whom you are going to share. If OrganizationNode
	// is passed in, PortfolioShare will be created for the node and its children
	// (when applies), and a PortfolioShareToken will be returned in the output
	// in order for the administrator to monitor the status of the PortfolioShare
	// creation process.
	OrganizationNode *OrganizationNode `type:"structure"`

	// The portfolio identifier.
	//
	// PortfolioId is a required field
	PortfolioId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreatePortfolioShareInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePortfolioShareInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePortfolioShareInput"}

	if s.PortfolioId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortfolioId"))
	}
	if s.PortfolioId != nil && len(*s.PortfolioId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PortfolioId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreatePortfolioShareOutput struct {
	_ struct{} `type:"structure"`

	// The portfolio share unique identifier. This will only be returned if portfolio
	// is shared to an organization node.
	PortfolioShareToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreatePortfolioShareOutput) String() string {
	return awsutil.Prettify(s)
}
