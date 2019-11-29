// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListActiveViolationsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return at one time.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token for the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The name of the Device Defender security profile for which violations are
	// listed.
	SecurityProfileName *string `location:"querystring" locationName:"securityProfileName" min:"1" type:"string"`

	// The name of the thing whose active violations are listed.
	ThingName *string `location:"querystring" locationName:"thingName" min:"1" type:"string"`
}

// String returns the string representation
func (s ListActiveViolationsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListActiveViolationsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListActiveViolationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.SecurityProfileName != nil && len(*s.SecurityProfileName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SecurityProfileName", 1))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListActiveViolationsOutput struct {
	_ struct{} `type:"structure"`

	// The list of active violations.
	ActiveViolations []ActiveViolation `locationName:"activeViolations" type:"list"`

	// A token that can be used to retrieve the next set of results, or null if
	// there are no additional results.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListActiveViolationsOutput) String() string {
	return awsutil.Prettify(s)
}