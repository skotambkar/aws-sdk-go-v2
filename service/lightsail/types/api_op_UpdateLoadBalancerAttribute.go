// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/enums"
)

type UpdateLoadBalancerAttributeInput struct {
	_ struct{} `type:"structure"`

	// The name of the attribute you want to update. Valid values are below.
	//
	// AttributeName is a required field
	AttributeName enums.LoadBalancerAttributeName `locationName:"attributeName" type:"string" required:"true" enum:"true"`

	// The value that you want to specify for the attribute name.
	//
	// AttributeValue is a required field
	AttributeValue *string `locationName:"attributeValue" min:"1" type:"string" required:"true"`

	// The name of the load balancer that you want to modify (e.g., my-load-balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateLoadBalancerAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateLoadBalancerAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateLoadBalancerAttributeInput"}
	if len(s.AttributeName) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("AttributeName"))
	}

	if s.AttributeValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttributeValue"))
	}
	if s.AttributeValue != nil && len(*s.AttributeValue) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AttributeValue", 1))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateLoadBalancerAttributeOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the API operations.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s UpdateLoadBalancerAttributeOutput) String() string {
	return awsutil.Prettify(s)
}