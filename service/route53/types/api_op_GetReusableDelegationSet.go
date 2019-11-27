// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to get information about a specified reusable delegation set.
type GetReusableDelegationSetInput struct {
	_ struct{} `type:"structure"`

	// The ID of the reusable delegation set that you want to get a list of name
	// servers for.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetReusableDelegationSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetReusableDelegationSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetReusableDelegationSetInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A complex type that contains the response to the GetReusableDelegationSet
// request.
type GetReusableDelegationSetOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about the reusable delegation set.
	//
	// DelegationSet is a required field
	DelegationSet *DelegationSet `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetReusableDelegationSetOutput) String() string {
	return awsutil.Prettify(s)
}
