// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddThingToThingGroupInput struct {
	_ struct{} `type:"structure"`

	// Override dynamic thing groups with static thing groups when 10-group limit
	// is reached. If a thing belongs to 10 thing groups, and one or more of those
	// groups are dynamic thing groups, adding a thing to a static group removes
	// the thing from the last dynamic group.
	OverrideDynamicGroups *bool `locationName:"overrideDynamicGroups" type:"boolean"`

	// The ARN of the thing to add to a group.
	ThingArn *string `locationName:"thingArn" type:"string"`

	// The ARN of the group to which you are adding a thing.
	ThingGroupArn *string `locationName:"thingGroupArn" type:"string"`

	// The name of the group to which you are adding a thing.
	ThingGroupName *string `locationName:"thingGroupName" min:"1" type:"string"`

	// The name of the thing to add to a group.
	ThingName *string `locationName:"thingName" min:"1" type:"string"`
}

// String returns the string representation
func (s AddThingToThingGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddThingToThingGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddThingToThingGroupInput"}
	if s.ThingGroupName != nil && len(*s.ThingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingGroupName", 1))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type AddThingToThingGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddThingToThingGroupOutput) String() string {
	return awsutil.Prettify(s)
}
