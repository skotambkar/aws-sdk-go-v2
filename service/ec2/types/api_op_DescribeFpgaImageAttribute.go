// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ec2/enums"
)

type DescribeFpgaImageAttributeInput struct {
	_ struct{} `type:"structure"`

	// The AFI attribute.
	//
	// Attribute is a required field
	Attribute enums.FpgaImageAttributeName `type:"string" required:"true" enum:"true"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the AFI.
	//
	// FpgaImageId is a required field
	FpgaImageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeFpgaImageAttributeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFpgaImageAttributeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFpgaImageAttributeInput"}
	if len(s.Attribute) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Attribute"))
	}

	if s.FpgaImageId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FpgaImageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeFpgaImageAttributeOutput struct {
	_ struct{} `type:"structure"`

	// Information about the attribute.
	FpgaImageAttribute *FpgaImageAttribute `locationName:"fpgaImageAttribute" type:"structure"`
}

// String returns the string representation
func (s DescribeFpgaImageAttributeOutput) String() string {
	return awsutil.Prettify(s)
}