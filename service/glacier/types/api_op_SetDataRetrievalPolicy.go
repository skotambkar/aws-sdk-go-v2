// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// SetDataRetrievalPolicy input.
type SetDataRetrievalPolicyInput struct {
	_ struct{} `type:"structure"`

	// The AccountId value is the AWS account ID. This value must match the AWS
	// account ID associated with the credentials used to sign the request. You
	// can either specify an AWS account ID or optionally a single '-' (hyphen),
	// in which case Amazon Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you specify your account ID, do
	// not include any hyphens ('-') in the ID.
	//
	// AccountId is a required field
	AccountId *string `location:"uri" locationName:"accountId" type:"string" required:"true"`

	// The data retrieval policy in JSON format.
	Policy *DataRetrievalPolicy `type:"structure"`
}

// String returns the string representation
func (s SetDataRetrievalPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetDataRetrievalPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetDataRetrievalPolicyInput"}

	if s.AccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SetDataRetrievalPolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetDataRetrievalPolicyOutput) String() string {
	return awsutil.Prettify(s)
}