// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeBrokerInstanceOptionsInput struct {
	_ struct{} `type:"structure"`

	EngineType *string `location:"querystring" locationName:"engineType" type:"string"`

	HostInstanceType *string `location:"querystring" locationName:"hostInstanceType" type:"string"`

	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeBrokerInstanceOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBrokerInstanceOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBrokerInstanceOptionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeBrokerInstanceOptionsOutput struct {
	_ struct{} `type:"structure"`

	BrokerInstanceOptions []BrokerInstanceOption `locationName:"brokerInstanceOptions" type:"list"`

	MaxResults *int64 `locationName:"maxResults" min:"5" type:"integer"`

	NextToken *string `locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s DescribeBrokerInstanceOptionsOutput) String() string {
	return awsutil.Prettify(s)
}