// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateSecurityGroupRuleDescriptionsEgressInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the security group. You must specify either the security group
	// ID or the security group name in the request. For security groups in a nondefault
	// VPC, you must specify the security group ID.
	GroupId *string `type:"string"`

	// [Default VPC] The name of the security group. You must specify either the
	// security group ID or the security group name in the request.
	GroupName *string `type:"string"`

	// The IP permissions for the security group rule.
	//
	// IpPermissions is a required field
	IpPermissions []IpPermission `locationNameList:"item" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateSecurityGroupRuleDescriptionsEgressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSecurityGroupRuleDescriptionsEgressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSecurityGroupRuleDescriptionsEgressInput"}

	if s.IpPermissions == nil {
		invalidParams.Add(aws.NewErrParamRequired("IpPermissions"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSecurityGroupRuleDescriptionsEgressOutput struct {
	_ struct{} `type:"structure"`

	// Returns true if the request succeeds; otherwise, returns an error.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s UpdateSecurityGroupRuleDescriptionsEgressOutput) String() string {
	return awsutil.Prettify(s)
}