// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListPartnerEventSourcesInput struct {
	_ struct{} `type:"structure"`

	// pecifying this limits the number of results returned by this operation. The
	// operation also returns a NextToken that you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int64 `min:"1" type:"integer"`

	// If you specify this, the results are limited to only those partner event
	// sources that start with the string you specify.
	//
	// NamePrefix is a required field
	NamePrefix *string `min:"1" type:"string" required:"true"`

	// The token returned by a previous call to this operation. Specifying this
	// retrieves the next set of results.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListPartnerEventSourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPartnerEventSourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListPartnerEventSourcesInput"}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if s.NamePrefix == nil {
		invalidParams.Add(aws.NewErrParamRequired("NamePrefix"))
	}
	if s.NamePrefix != nil && len(*s.NamePrefix) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NamePrefix", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListPartnerEventSourcesOutput struct {
	_ struct{} `type:"structure"`

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string `min:"1" type:"string"`

	// The list of partner event sources returned by the operation.
	PartnerEventSources []PartnerEventSource `type:"list"`
}

// String returns the string representation
func (s ListPartnerEventSourcesOutput) String() string {
	return awsutil.Prettify(s)
}
