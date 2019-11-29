// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateIndexingConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Thing group indexing configuration.
	ThingGroupIndexingConfiguration *ThingGroupIndexingConfiguration `locationName:"thingGroupIndexingConfiguration" type:"structure"`

	// Thing indexing configuration.
	ThingIndexingConfiguration *ThingIndexingConfiguration `locationName:"thingIndexingConfiguration" type:"structure"`
}

// String returns the string representation
func (s UpdateIndexingConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateIndexingConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateIndexingConfigurationInput"}
	if s.ThingGroupIndexingConfiguration != nil {
		if err := s.ThingGroupIndexingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ThingGroupIndexingConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.ThingIndexingConfiguration != nil {
		if err := s.ThingIndexingConfiguration.Validate(); err != nil {
			invalidParams.AddNested("ThingIndexingConfiguration", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateIndexingConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateIndexingConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}
