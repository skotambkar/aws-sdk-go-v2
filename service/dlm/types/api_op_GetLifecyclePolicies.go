// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/dlm/enums"
)

type GetLifecyclePoliciesInput struct {
	_ struct{} `type:"structure"`

	// The identifiers of the data lifecycle policies.
	PolicyIds []string `location:"querystring" locationName:"policyIds" type:"list"`

	// The resource type.
	ResourceTypes []enums.ResourceTypeValues `location:"querystring" locationName:"resourceTypes" min:"1" type:"list"`

	// The activation state.
	State enums.GettablePolicyStateValues `location:"querystring" locationName:"state" type:"string" enum:"true"`

	// The tags to add to objects created by the policy.
	//
	// Tags are strings in the format key=value.
	//
	// These user-defined tags are added in addition to the AWS-added lifecycle
	// tags.
	TagsToAdd []string `location:"querystring" locationName:"tagsToAdd" type:"list"`

	// The target tag for a policy.
	//
	// Tags are strings in the format key=value.
	TargetTags []string `location:"querystring" locationName:"targetTags" min:"1" type:"list"`
}

// String returns the string representation
func (s GetLifecyclePoliciesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetLifecyclePoliciesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetLifecyclePoliciesInput"}
	if s.ResourceTypes != nil && len(s.ResourceTypes) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ResourceTypes", 1))
	}
	if s.TargetTags != nil && len(s.TargetTags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TargetTags", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetLifecyclePoliciesOutput struct {
	_ struct{} `type:"structure"`

	// Summary information about the lifecycle policies.
	Policies []LifecyclePolicySummary `type:"list"`
}

// String returns the string representation
func (s GetLifecyclePoliciesOutput) String() string {
	return awsutil.Prettify(s)
}
