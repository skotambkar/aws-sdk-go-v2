// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AssociateAdminAccountInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID to associate with AWS Firewall Manager as the AWS Firewall
	// Manager administrator account. This can be an AWS Organizations master account
	// or a member account. For more information about AWS Organizations and master
	// accounts, see Managing the AWS Accounts in Your Organization (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts.html).
	//
	// AdminAccount is a required field
	AdminAccount *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateAdminAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateAdminAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateAdminAccountInput"}

	if s.AdminAccount == nil {
		invalidParams.Add(aws.NewErrParamRequired("AdminAccount"))
	}
	if s.AdminAccount != nil && len(*s.AdminAccount) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AdminAccount", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AssociateAdminAccountOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AssociateAdminAccountOutput) String() string {
	return awsutil.Prettify(s)
}