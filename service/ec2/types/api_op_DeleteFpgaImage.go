// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFpgaImageInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the AFI.
	//
	// FpgaImageId is a required field
	FpgaImageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFpgaImageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFpgaImageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFpgaImageInput"}

	if s.FpgaImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FpgaImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFpgaImageOutput struct {
	_ struct{} `type:"structure"`

	// Is true if the request succeeds, and an error otherwise.
	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteFpgaImageOutput) String() string {
	return awsutil.Prettify(s)
}
