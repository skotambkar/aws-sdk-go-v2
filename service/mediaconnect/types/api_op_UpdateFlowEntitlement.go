// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The updates that you want to make to a specific entitlement.
type UpdateFlowEntitlementInput struct {
	_ struct{} `type:"structure"`

	// A description of the entitlement. This description appears only on the AWS
	// Elemental MediaConnect console and will not be seen by the subscriber or
	// end user.
	Description *string `locationName:"description" type:"string"`

	// The type of encryption that will be used on the output associated with this
	// entitlement.
	Encryption *UpdateEncryption `locationName:"encryption" type:"structure"`

	// EntitlementArn is a required field
	EntitlementArn *string `location:"uri" locationName:"entitlementArn" type:"string" required:"true"`

	// FlowArn is a required field
	FlowArn *string `location:"uri" locationName:"flowArn" type:"string" required:"true"`

	// The AWS account IDs that you want to share your content with. The receiving
	// accounts (subscribers) will be allowed to create their own flow using your
	// content as the source.
	Subscribers []string `locationName:"subscribers" type:"list"`
}

// String returns the string representation
func (s UpdateFlowEntitlementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateFlowEntitlementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateFlowEntitlementInput"}

	if s.EntitlementArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EntitlementArn"))
	}

	if s.FlowArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("FlowArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a successful UpdateFlowEntitlement request. The response includes
// the ARN of the flow that was updated and the updated entitlement configuration.
type UpdateFlowEntitlementOutput struct {
	_ struct{} `type:"structure"`

	// The settings for a flow entitlement.
	Entitlement *Entitlement `locationName:"entitlement" type:"structure"`

	// The ARN of the flow that this entitlement was granted on.
	FlowArn *string `locationName:"flowArn" type:"string"`
}

// String returns the string representation
func (s UpdateFlowEntitlementOutput) String() string {
	return awsutil.Prettify(s)
}
