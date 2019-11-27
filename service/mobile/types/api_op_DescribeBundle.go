// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure to request the details of a specific bundle.
type DescribeBundleInput struct {
	_ struct{} `type:"structure"`

	// Unique bundle identifier.
	//
	// BundleId is a required field
	BundleId *string `location:"uri" locationName:"bundleId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBundleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBundleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeBundleInput"}

	if s.BundleId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BundleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Result structure contains the details of the bundle.
type DescribeBundleOutput struct {
	_ struct{} `type:"structure"`

	// The details of the bundle.
	Details *BundleDetails `locationName:"details" type:"structure"`
}

// String returns the string representation
func (s DescribeBundleOutput) String() string {
	return awsutil.Prettify(s)
}
