// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateGeoMatchSetInput struct {
	_ struct{} `type:"structure"`

	// The value returned by the most recent call to GetChangeToken.
	//
	// ChangeToken is a required field
	ChangeToken *string `min:"1" type:"string" required:"true"`

	// A friendly name or description of the GeoMatchSet. You can't change Name
	// after you create the GeoMatchSet.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateGeoMatchSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGeoMatchSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGeoMatchSetInput"}

	if s.ChangeToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ChangeToken"))
	}
	if s.ChangeToken != nil && len(*s.ChangeToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ChangeToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGeoMatchSetOutput struct {
	_ struct{} `type:"structure"`

	// The ChangeToken that you used to submit the CreateGeoMatchSet request. You
	// can also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string `min:"1" type:"string"`

	// The GeoMatchSet returned in the CreateGeoMatchSet response. The GeoMatchSet
	// contains no GeoMatchConstraints.
	GeoMatchSet *GeoMatchSet `type:"structure"`
}

// String returns the string representation
func (s CreateGeoMatchSetOutput) String() string {
	return awsutil.Prettify(s)
}