// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/wafregional/enums"
)

type ListResourcesForWebACLInput struct {
	_ struct{} `type:"structure"`

	// The type of resource to list, either an application load balancer or Amazon
	// API Gateway.
	ResourceType enums.ResourceType `type:"string" enum:"true"`

	// The unique identifier (ID) of the web ACL for which to list the associated
	// resources.
	//
	// WebACLId is a required field
	WebACLId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ListResourcesForWebACLInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResourcesForWebACLInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResourcesForWebACLInput"}

	if s.WebACLId == nil {
		invalidParams.Add(aws.NewErrParamRequired("WebACLId"))
	}
	if s.WebACLId != nil && len(*s.WebACLId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("WebACLId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListResourcesForWebACLOutput struct {
	_ struct{} `type:"structure"`

	// An array of ARNs (Amazon Resource Names) of the resources associated with
	// the specified web ACL. An array with zero elements is returned if there are
	// no resources associated with the web ACL.
	ResourceArns []string `type:"list"`
}

// String returns the string representation
func (s ListResourcesForWebACLOutput) String() string {
	return awsutil.Prettify(s)
}