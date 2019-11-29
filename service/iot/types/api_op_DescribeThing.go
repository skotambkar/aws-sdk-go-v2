// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the DescribeThing operation.
type DescribeThingInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeThingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeThingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeThingInput"}

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

// The output from the DescribeThing operation.
type DescribeThingOutput struct {
	_ struct{} `type:"structure"`

	// The thing attributes.
	Attributes map[string]string `locationName:"attributes" type:"map"`

	// The name of the billing group the thing belongs to.
	BillingGroupName *string `locationName:"billingGroupName" min:"1" type:"string"`

	// The default client ID.
	DefaultClientId *string `locationName:"defaultClientId" type:"string"`

	// The ARN of the thing to describe.
	ThingArn *string `locationName:"thingArn" type:"string"`

	// The ID of the thing to describe.
	ThingId *string `locationName:"thingId" type:"string"`

	// The name of the thing.
	ThingName *string `locationName:"thingName" min:"1" type:"string"`

	// The thing type name.
	ThingTypeName *string `locationName:"thingTypeName" min:"1" type:"string"`

	// The current version of the thing record in the registry.
	//
	// To avoid unintentional changes to the information in the registry, you can
	// pass the version information in the expectedVersion parameter of the UpdateThing
	// and DeleteThing calls.
	Version *int64 `locationName:"version" type:"long"`
}

// String returns the string representation
func (s DescribeThingOutput) String() string {
	return awsutil.Prettify(s)
}