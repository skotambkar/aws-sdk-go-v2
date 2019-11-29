// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteTrafficMirrorFilterRuleInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the Traffic Mirror rule.
	//
	// TrafficMirrorFilterRuleId is a required field
	TrafficMirrorFilterRuleId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTrafficMirrorFilterRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTrafficMirrorFilterRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTrafficMirrorFilterRuleInput"}

	if s.TrafficMirrorFilterRuleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrafficMirrorFilterRuleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteTrafficMirrorFilterRuleOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the deleted Traffic Mirror rule.
	TrafficMirrorFilterRuleId *string `locationName:"trafficMirrorFilterRuleId" type:"string"`
}

// String returns the string representation
func (s DeleteTrafficMirrorFilterRuleOutput) String() string {
	return awsutil.Prettify(s)
}