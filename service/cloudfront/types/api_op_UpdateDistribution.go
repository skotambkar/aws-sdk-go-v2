// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request to update a distribution.
type UpdateDistributionInput struct {
	_ struct{} `type:"structure" payload:"DistributionConfig"`

	// The distribution's configuration information.
	//
	// DistributionConfig is a required field
	DistributionConfig *DistributionConfig `locationName:"DistributionConfig" type:"structure" required:"true" xmlURI:"http://cloudfront.amazonaws.com/doc/2019-03-26/"`

	// The distribution's id.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header that you received when retrieving the distribution's
	// configuration. For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`
}

// String returns the string representation
func (s UpdateDistributionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDistributionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDistributionInput"}

	if s.DistributionConfig == nil {
		invalidParams.Add(aws.NewErrParamRequired("DistributionConfig"))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.DistributionConfig != nil {
		if err := s.DistributionConfig.Validate(); err != nil {
			invalidParams.AddNested("DistributionConfig", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The returned result of the corresponding request.
type UpdateDistributionOutput struct {
	_ struct{} `type:"structure" payload:"Distribution"`

	// The distribution's information.
	Distribution *Distribution `type:"structure"`

	// The current version of the configuration. For example: E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`
}

// String returns the string representation
func (s UpdateDistributionOutput) String() string {
	return awsutil.Prettify(s)
}