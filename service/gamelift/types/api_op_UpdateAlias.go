// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for a request action.
type UpdateAliasInput struct {
	_ struct{} `type:"structure"`

	// Unique identifier for a fleet alias. Specify the alias you want to update.
	//
	// AliasId is a required field
	AliasId *string `type:"string" required:"true"`

	// Human-readable description of an alias.
	Description *string `min:"1" type:"string"`

	// Descriptive label that is associated with an alias. Alias names do not need
	// to be unique.
	Name *string `min:"1" type:"string"`

	// Object that specifies the fleet and routing type to use for the alias.
	RoutingStrategy *RoutingStrategy `type:"structure"`
}

// String returns the string representation
func (s UpdateAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAliasInput"}

	if s.AliasId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasId"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the returned data in response to a request action.
type UpdateAliasOutput struct {
	_ struct{} `type:"structure"`

	// Object that contains the updated alias configuration.
	Alias *Alias `type:"structure"`
}

// String returns the string representation
func (s UpdateAliasOutput) String() string {
	return awsutil.Prettify(s)
}