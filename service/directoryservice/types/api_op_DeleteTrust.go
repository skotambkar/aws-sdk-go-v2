// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Deletes the local side of an existing trust relationship between the AWS
// Managed Microsoft AD directory and the external domain.
type DeleteTrustInput struct {
	_ struct{} `type:"structure"`

	// Delete a conditional forwarder as part of a DeleteTrustRequest.
	DeleteAssociatedConditionalForwarder *bool `type:"boolean"`

	// The Trust ID of the trust relationship to be deleted.
	//
	// TrustId is a required field
	TrustId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTrustInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTrustInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTrustInput"}

	if s.TrustId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrustId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DeleteTrust request.
type DeleteTrustOutput struct {
	_ struct{} `type:"structure"`

	// The Trust ID of the trust relationship that was deleted.
	TrustId *string `type:"string"`
}

// String returns the string representation
func (s DeleteTrustOutput) String() string {
	return awsutil.Prettify(s)
}
