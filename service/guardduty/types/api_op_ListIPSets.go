// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListIPSetsInput struct {
	_ struct{} `type:"structure"`

	// The unique ID of the detector the ipSet is associated with.
	//
	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" min:"1" type:"string" required:"true"`

	// You can use this parameter to indicate the maximum number of items you want
	// in the response. The default value is 50. The maximum value is 50.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the list action. For subsequent calls
	// to the action fill nextToken in the request with the value of NextToken from
	// the previous response to continue listing data.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListIPSetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListIPSetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListIPSetsInput"}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListIPSetsOutput struct {
	_ struct{} `type:"structure"`

	// The IDs of the IPSet resources.
	//
	// IpSetIds is a required field
	IpSetIds []string `locationName:"ipSetIds" type:"list" required:"true"`

	// Pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListIPSetsOutput) String() string {
	return awsutil.Prettify(s)
}
