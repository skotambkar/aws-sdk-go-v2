// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetEventSelectorsInput struct {
	_ struct{} `type:"structure"`

	// Specifies the name of the trail or trail ARN. If you specify a trail name,
	// the string must meet the following requirements:
	//
	//    * Contain only ASCII letters (a-z, A-Z), numbers (0-9), periods (.), underscores
	//    (_), or dashes (-)
	//
	//    * Start with a letter or number, and end with a letter or number
	//
	//    * Be between 3 and 128 characters
	//
	//    * Have no adjacent periods, underscores or dashes. Names like my-_namespace
	//    and my--namespace are not valid.
	//
	//    * Not be in IP address format (for example, 192.168.5.4)
	//
	// If you specify a trail ARN, it must be in the format:
	//
	// arn:aws:cloudtrail:us-east-2:123456789012:trail/MyTrail
	//
	// TrailName is a required field
	TrailName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetEventSelectorsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEventSelectorsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEventSelectorsInput"}

	if s.TrailName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrailName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetEventSelectorsOutput struct {
	_ struct{} `type:"structure"`

	// The event selectors that are configured for the trail.
	EventSelectors []EventSelector `type:"list"`

	// The specified trail ARN that has the event selectors.
	TrailARN *string `type:"string"`
}

// String returns the string representation
func (s GetEventSelectorsOutput) String() string {
	return awsutil.Prettify(s)
}