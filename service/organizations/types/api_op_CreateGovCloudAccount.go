// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/organizations/enums"
)

type CreateGovCloudAccountInput struct {
	_ struct{} `type:"structure"`

	// The friendly name of the member account.
	//
	// AccountName is a required field
	AccountName *string `min:"1" type:"string" required:"true" sensitive:"true"`

	// The email address of the owner to assign to the new member account in the
	// commercial Region. This email address must not already be associated with
	// another AWS account. You must use a valid email address to complete account
	// creation. You can't access the root user of the account or remove an account
	// that was created with an invalid email address. Like all request parameters
	// for CreateGovCloudAccount, the request for the email address for the AWS
	// GovCloud (US) account originates from the commercial Region, not from the
	// AWS GovCloud (US) Region.
	//
	// Email is a required field
	Email *string `min:"6" type:"string" required:"true" sensitive:"true"`

	// If set to ALLOW, the new linked account in the commercial Region enables
	// IAM users to access account billing information if they have the required
	// permissions. If set to DENY, only the root user of the new account can access
	// account billing information. For more information, see Activating Access
	// to the Billing and Cost Management Console (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/grantaccess.html#ControllingAccessWebsite-Activate)
	// in the AWS Billing and Cost Management User Guide.
	//
	// If you don't specify this parameter, the value defaults to ALLOW, and IAM
	// users and roles with the required permissions can access billing information
	// for the new account.
	IamUserAccessToBilling enums.IAMUserAccessToBilling `type:"string" enum:"true"`

	// (Optional)
	//
	// The name of an IAM role that AWS Organizations automatically preconfigures
	// in the new member accounts in both the AWS GovCloud (US) Region and in the
	// commercial Region. This role trusts the master account, allowing users in
	// the master account to assume the role, as permitted by the master account
	// administrator. The role has administrator permissions in the new member account.
	//
	// If you don't specify this parameter, the role name defaults to OrganizationAccountAccessRole.
	//
	// For more information about how to use this role to access the member account,
	// see Accessing and Administering the Member Accounts in Your Organization
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_accounts_access.html#orgs_manage_accounts_create-cross-account-role)
	// in the AWS Organizations User Guide and steps 2 and 3 in Tutorial: Delegate
	// Access Across AWS Accounts Using IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/tutorial_cross-account-with-roles.html)
	// in the IAM User Guide.
	//
	// The regex pattern (http://wikipedia.org/wiki/regex) that is used to validate
	// this parameter is a string of characters that can consist of uppercase letters,
	// lowercase letters, digits with no spaces, and any of the following characters:
	// =,.@-
	RoleName *string `type:"string"`
}

// String returns the string representation
func (s CreateGovCloudAccountInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGovCloudAccountInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGovCloudAccountInput"}

	if s.AccountName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountName"))
	}
	if s.AccountName != nil && len(*s.AccountName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AccountName", 1))
	}

	if s.Email == nil {
		invalidParams.Add(aws.NewErrParamRequired("Email"))
	}
	if s.Email != nil && len(*s.Email) < 6 {
		invalidParams.Add(aws.NewErrParamMinLen("Email", 6))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGovCloudAccountOutput struct {
	_ struct{} `type:"structure"`

	// Contains the status about a CreateAccount or CreateGovCloudAccount request
	// to create an AWS account or an AWS GovCloud (US) account in an organization.
	CreateAccountStatus *CreateAccountStatus `type:"structure"`
}

// String returns the string representation
func (s CreateGovCloudAccountOutput) String() string {
	return awsutil.Prettify(s)
}
