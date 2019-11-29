// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Deletes a origin access identity.
type DeleteCloudFrontOriginAccessIdentityInput struct {
	_ struct{} `type:"structure"`

	// The origin access identity's ID.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header you received from a previous GET or PUT request.
	// For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`
}

// String returns the string representation
func (s DeleteCloudFrontOriginAccessIdentityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCloudFrontOriginAccessIdentityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteCloudFrontOriginAccessIdentityInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteCloudFrontOriginAccessIdentityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteCloudFrontOriginAccessIdentityOutput) String() string {
	return awsutil.Prettify(s)
}