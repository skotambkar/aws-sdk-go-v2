// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetTrailInput struct {
	_ struct{} `type:"structure"`

	// The name or the Amazon Resource Name (ARN) of the trail for which you want
	// to retrieve settings information.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetTrailInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetTrailInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetTrailInput"}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetTrailOutput struct {
	_ struct{} `type:"structure"`

	// The settings for a trail.
	Trail *Trail `type:"structure"`
}

// String returns the string representation
func (s GetTrailOutput) String() string {
	return awsutil.Prettify(s)
}