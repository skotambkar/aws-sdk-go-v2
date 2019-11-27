// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListProvisionedCapacityInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID of the account that owns the vault. You can either specify
	// an AWS account ID or optionally a single '-' (hyphen), in which case Amazon
	// S3 Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you use an account ID, don't include any hyphens ('-')
	// in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListProvisionedCapacityInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListProvisionedCapacityInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListProvisionedCapacityInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListProvisionedCapacityOutput struct {
	_ struct{} `type:"structure"`

	// The response body contains the following JSON fields.
	ProvisionedCapacityList []ProvisionedCapacityDescription `type:"list"`
}

// String returns the string representation
func (s ListProvisionedCapacityOutput) String() string {
	return awsutil.Prettify(s)
}
