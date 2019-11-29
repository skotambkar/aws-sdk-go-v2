// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeEffectivePatchesForPatchBaselineInput struct {
	_ struct{} `type:"structure"`

	// The ID of the patch baseline to retrieve the effective patches for.
	//
	// BaselineId is a required field
	BaselineId *string `min:"20" type:"string" required:"true"`

	// The maximum number of patches to return (per page).
	MaxResults *int64 `min:"1" type:"integer"`

	// The token for the next set of items to return. (You received this token from
	// a previous call.)
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEffectivePatchesForPatchBaselineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEffectivePatchesForPatchBaselineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeEffectivePatchesForPatchBaselineInput"}

	if s.BaselineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BaselineId"))
	}
	if s.BaselineId != nil && len(*s.BaselineId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("BaselineId", 20))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeEffectivePatchesForPatchBaselineOutput struct {
	_ struct{} `type:"structure"`

	// An array of patches and patch status.
	EffectivePatches []EffectivePatch `type:"list"`

	// The token to use when requesting the next set of items. If there are no additional
	// items to return, the string is empty.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s DescribeEffectivePatchesForPatchBaselineOutput) String() string {
	return awsutil.Prettify(s)
}
