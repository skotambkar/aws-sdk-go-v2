// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSmartHomeAppliancesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of appliances to be returned, per paginated calls.
	MaxResults *int64 `min:"1" type:"integer"`

	// The tokens used for pagination.
	NextToken *string `min:"1" type:"string"`

	// The room that the appliances are associated with.
	//
	// RoomArn is a required field
	RoomArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListSmartHomeAppliancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSmartHomeAppliancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSmartHomeAppliancesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if s.RoomArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("RoomArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSmartHomeAppliancesOutput struct {
	_ struct{} `type:"structure"`

	// The tokens used for pagination.
	NextToken *string `min:"1" type:"string"`

	// The smart home appliances.
	SmartHomeAppliances []SmartHomeAppliance `type:"list"`
}

// String returns the string representation
func (s ListSmartHomeAppliancesOutput) String() string {
	return awsutil.Prettify(s)
}
