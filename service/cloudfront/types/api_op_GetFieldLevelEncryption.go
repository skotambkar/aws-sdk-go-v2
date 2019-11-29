// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetFieldLevelEncryptionInput struct {
	_ struct{} `type:"structure"`

	// Request the ID for the field-level encryption configuration information.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetFieldLevelEncryptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFieldLevelEncryptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFieldLevelEncryptionInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetFieldLevelEncryptionOutput struct {
	_ struct{} `type:"structure" payload:"FieldLevelEncryption"`

	// The current version of the field level encryption configuration. For example:
	// E2QWRUHAPOMQZL.
	ETag *string `location:"header" locationName:"ETag" type:"string"`

	// Return the field-level encryption configuration information.
	FieldLevelEncryption *FieldLevelEncryption `type:"structure"`
}

// String returns the string representation
func (s GetFieldLevelEncryptionOutput) String() string {
	return awsutil.Prettify(s)
}