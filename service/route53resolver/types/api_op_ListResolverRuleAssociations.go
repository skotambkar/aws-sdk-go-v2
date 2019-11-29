// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListResolverRuleAssociationsInput struct {
	_ struct{} `type:"structure"`

	// An optional specification to return a subset of resolver rules, such as resolver
	// rules that are associated with the same VPC ID.
	//
	// If you submit a second or subsequent ListResolverRuleAssociations request
	// and specify the NextToken parameter, you must use the same values for Filters,
	// if any, as in the previous request.
	Filters []Filter `type:"list"`

	// The maximum number of rule associations that you want to return in the response
	// to a ListResolverRuleAssociations request. If you don't specify a value for
	// MaxResults, Resolver returns up to 100 rule associations.
	MaxResults *int64 `min:"1" type:"integer"`

	// For the first ListResolverRuleAssociation request, omit this value.
	//
	// If you have more than MaxResults rule associations, you can submit another
	// ListResolverRuleAssociation request to get the next group of rule associations.
	// In the next request, specify the value of NextToken from the previous response.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListResolverRuleAssociationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListResolverRuleAssociationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListResolverRuleAssociationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListResolverRuleAssociationsOutput struct {
	_ struct{} `type:"structure"`

	// The value that you specified for MaxResults in the request.
	MaxResults *int64 `min:"1" type:"integer"`

	// If more than MaxResults rule associations match the specified criteria, you
	// can submit another ListResolverRuleAssociation request to get the next group
	// of results. In the next request, specify the value of NextToken from the
	// previous response.
	NextToken *string `type:"string"`

	// The associations that were created between resolver rules and VPCs using
	// the current AWS account, and that match the specified filters, if any.
	ResolverRuleAssociations []ResolverRuleAssociation `type:"list"`
}

// String returns the string representation
func (s ListResolverRuleAssociationsOutput) String() string {
	return awsutil.Prettify(s)
}