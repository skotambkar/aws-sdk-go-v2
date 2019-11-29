// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFleetInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// Fleet is a required field
	Fleet *string `locationName:"fleet" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteFleetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFleetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFleetInput"}

	if s.Fleet == nil {
		invalidParams.Add(aws.NewErrParamRequired("Fleet"))
	}
	if s.Fleet != nil && len(*s.Fleet) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fleet", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFleetOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFleetOutput) String() string {
	return awsutil.Prettify(s)
}