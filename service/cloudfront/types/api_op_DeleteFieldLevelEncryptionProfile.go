// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteFieldLevelEncryptionProfileInput struct {
	_ struct{} `type:"structure"`

	// Request the ID of the profile you want to delete from CloudFront.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`

	// The value of the ETag header that you received when retrieving the profile
	// to delete. For example: E2QWRUHAPOMQZL.
	IfMatch *string `location:"header" locationName:"If-Match" type:"string"`
}

// String returns the string representation
func (s DeleteFieldLevelEncryptionProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteFieldLevelEncryptionProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteFieldLevelEncryptionProfileInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteFieldLevelEncryptionProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteFieldLevelEncryptionProfileOutput) String() string {
	return awsutil.Prettify(s)
}