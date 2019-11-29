// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeHomeRegionControlsInput struct {
	_ struct{} `type:"structure"`

	// The ControlID is a unique identifier string of your HomeRegionControl object.
	ControlId *string `min:"1" type:"string"`

	// The name of the home region you'd like to view.
	HomeRegion *string `min:"1" type:"string"`

	// The maximum number of filtering results to display per page.
	MaxResults *int64 `min:"1" type:"integer"`

	// If a NextToken was returned by a previous call, more results are available.
	// To retrieve the next page of results, make the call again using the returned
	// token in NextToken.
	NextToken *string `type:"string"`

	// The target parameter specifies the identifier to which the home region is
	// applied, which is always of type ACCOUNT. It applies the home region to the
	// current ACCOUNT.
	Target *Target `type:"structure"`
}

// String returns the string representation
func (s DescribeHomeRegionControlsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeHomeRegionControlsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeHomeRegionControlsInput"}
	if s.ControlId != nil && len(*s.ControlId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ControlId", 1))
	}
	if s.HomeRegion != nil && len(*s.HomeRegion) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("HomeRegion", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.Target != nil {
		if err := s.Target.Validate(); err != nil {
			invalidParams.AddNested("Target", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeHomeRegionControlsOutput struct {
	_ struct{} `type:"structure"`

	// An array that contains your HomeRegionControl objects.
	HomeRegionControls []HomeRegionControl `type:"list"`

	// If a NextToken was returned by a previous call, more results are available.
	// To retrieve the next page of results, make the call again using the returned
	// token in NextToken.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeHomeRegionControlsOutput) String() string {
	return awsutil.Prettify(s)
}