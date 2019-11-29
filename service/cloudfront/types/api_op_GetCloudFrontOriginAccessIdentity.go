// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// The request to get an origin access identity's information.
type GetCloudFrontOriginAccessIdentityInput struct {
	_ struct{} `type:"structure"`

	// The identity's ID.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCloudFrontOriginAccessIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCloudFrontOriginAccessIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCloudFrontOriginAccessIdentityInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The returned result of the corresponding request.
type GetCloudFrontOriginAccessIdentityOutput struct {
	_ struct{} `type:"structure" payload:"CloudFrontOriginAccessIdentity"`

	// The origin access identity's information.
	CloudFrontOriginAccessIdentity *CloudFrontOriginAccessIdentity `type:"structure"`

	// The current version of the origin access identity's information. For example:
	// E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`
}

// String returns the string representation
func (s GetCloudFrontOriginAccessIdentityOutput) String() string {
	return awsutil.Prettify(s)
}