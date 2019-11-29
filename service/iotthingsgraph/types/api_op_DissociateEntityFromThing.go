// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/enums"
)

type DissociateEntityFromThingInput struct {
	_ struct{} `type:"structure"`

	// The entity type from which to disassociate the thing.
	//
	// EntityType is a required field
	EntityType enums.EntityType `locationName:"entityType" type:"string" required:"true" enum:"true"`

	// The name of the thing to disassociate.
	//
	// ThingName is a required field
	ThingName *string `locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DissociateEntityFromThingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DissociateEntityFromThingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DissociateEntityFromThingInput"}
	if len(s.EntityType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("EntityType"))
	}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DissociateEntityFromThingOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DissociateEntityFromThingOutput) String() string {
	return awsutil.Prettify(s)
}