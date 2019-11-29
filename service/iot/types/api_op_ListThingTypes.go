// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The input for the ListThingTypes operation.
type ListThingTypesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return in this operation.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The name of the thing type.
	ThingTypeName *string `location:"querystring" locationName:"thingTypeName" min:"1" type:"string"`
}

// String returns the string representation
func (s ListThingTypesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListThingTypesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListThingTypesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.ThingTypeName != nil && len(*s.ThingTypeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ThingTypeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The output for the ListThingTypes operation.
type ListThingTypesOutput struct {
	_ struct{} `type:"structure"`

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The thing types.
	ThingTypes []ThingTypeDefinition `locationName:"thingTypes" type:"list"`
}

// String returns the string representation
func (s ListThingTypesOutput) String() string {
	return awsutil.Prettify(s)
}