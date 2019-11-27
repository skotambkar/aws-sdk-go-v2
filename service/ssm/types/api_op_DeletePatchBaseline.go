// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeletePatchBaselineInput struct {
	_ struct{} `type:"structure"`

	// The ID of the patch baseline to delete.
	//
	// BaselineId is a required field
	BaselineId *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePatchBaselineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePatchBaselineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePatchBaselineInput"}

	if s.BaselineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BaselineId"))
	}
	if s.BaselineId != nil && len(*s.BaselineId) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("BaselineId", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePatchBaselineOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the deleted patch baseline.
	BaselineId *string `min:"20" type:"string"`
}

// String returns the string representation
func (s DeletePatchBaselineOutput) String() string {
	return awsutil.Prettify(s)
}
