// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterUsageInput struct {
	_ struct{} `type:"structure"`

	// (Optional) To scope down the registration to a specific running software
	// instance and guard against replay attacks.
	Nonce *string `type:"string"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// Public Key Version provided by AWS Marketplace
	//
	// PublicKeyVersion is a required field
	PublicKeyVersion *int64 `min:"1" type:"integer" required:"true"`
}

// String returns the string representation
func (s RegisterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterUsageInput"}

	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductCode", 1))
	}

	if s.PublicKeyVersion == nil {
		invalidParams.Add(aws.NewErrParamRequired("PublicKeyVersion"))
	}
	if s.PublicKeyVersion != nil && *s.PublicKeyVersion < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PublicKeyVersion", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RegisterUsageOutput struct {
	_ struct{} `type:"structure"`

	// (Optional) Only included when public key version has expired
	PublicKeyRotationTimestamp *time.Time `type:"timestamp"`

	// JWT Token
	Signature *string `type:"string"`
}

// String returns the string representation
func (s RegisterUsageOutput) String() string {
	return awsutil.Prettify(s)
}