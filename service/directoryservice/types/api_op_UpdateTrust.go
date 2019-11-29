// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/directoryservice/enums"
)

type UpdateTrustInput struct {
	_ struct{} `type:"structure"`

	// Updates selective authentication for the trust.
	SelectiveAuth enums.SelectiveAuth `type:"string" enum:"true"`

	// Identifier of the trust relationship.
	//
	// TrustId is a required field
	TrustId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTrustInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTrustInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTrustInput"}

	if s.TrustId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TrustId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateTrustOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request identifier.
	RequestId *string `type:"string"`

	// Identifier of the trust relationship.
	TrustId *string `type:"string"`
}

// String returns the string representation
func (s UpdateTrustOutput) String() string {
	return awsutil.Prettify(s)
}
