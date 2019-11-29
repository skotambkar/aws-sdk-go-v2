// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type StopGameSessionPlacementInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a game session placement to cancel.
	//
	// PlacementId is a required field
	PlacementId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s StopGameSessionPlacementInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StopGameSessionPlacementInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StopGameSessionPlacementInput"}

	if s.PlacementId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PlacementId"))
	}
	if s.PlacementId != nil && len(*s.PlacementId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("PlacementId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type StopGameSessionPlacementOutput struct {
	_ struct{} `type:"structure"`

	// Object that describes the canceled game session placement, with CANCELLED
	// status and an end time stamp.
	GameSessionPlacement *GameSessionPlacement `type:"structure"`
}

// String returns the string representation
func (s StopGameSessionPlacementOutput) String() string {
	return awsutil.Prettify(s)
}
