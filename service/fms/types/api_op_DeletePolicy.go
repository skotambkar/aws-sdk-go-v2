// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeletePolicyInput struct {
	_ struct{} `type:"structure"`

	// If True, the request performs cleanup according to the policy type.
	//
	// For AWS WAF and Shield Advanced policies, the cleanup does the following:
	//
	//    * Deletes rule groups created by AWS Firewall Manager
	//
	//    * Removes web ACLs from in-scope resources
	//
	//    * Deletes web ACLs that contain no rules or rule groups
	//
	// For security group policies, the cleanup does the following for each security
	// group in the policy:
	//
	//    * Disassociates the security group from in-scope resources
	//
	//    * Deletes the security group if it was created through Firewall Manager
	//    and if it's no longer associated with any resources through another policy
	//
	// After the cleanup, in-scope resources are no longer protected by web ACLs
	// in this policy. Protection of out-of-scope resources remains unchanged. Scope
	// is determined by tags that you create and accounts that you associate with
	// the policy. When creating the policy, if you specify that only resources
	// in specific accounts or with specific tags are in scope of the policy, those
	// accounts and resources are handled by the policy. All others are out of scope.
	// If you don't specify tags or accounts, all resources are in scope.
	DeleteAllPolicyResources *bool `type:"boolean"`

	// The ID of the policy that you want to delete. PolicyId is returned by PutPolicy
	// and by ListPolicies.
	//
	// PolicyId is a required field
	PolicyId *string `min:"36" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePolicyInput"}

	if s.PolicyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyId"))
	}
	if s.PolicyId != nil && len(*s.PolicyId) < 36 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyId", 36))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePolicyOutput) String() string {
	return awsutil.Prettify(s)
}