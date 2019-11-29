// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the DescribeAvailabilityOptions operation.
// Specifies the name of the domain you want to describe. To show the active
// configuration and exclude any pending changes, set the Deployed option to
// true.
type DescribeAvailabilityOptionsInput struct {
	_ struct{} `type:"structure"`

	// Whether to display the deployed configuration (true) or include any pending
	// changes (false). Defaults to false.
	Deployed *bool `type:"boolean"`

	// The name of the domain you want to describe.
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAvailabilityOptionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAvailabilityOptionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAvailabilityOptionsInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a DescribeAvailabilityOptions request. Indicates whether or
// not the Multi-AZ option is enabled for the domain specified in the request.
type DescribeAvailabilityOptionsOutput struct {
	_ struct{} `type:"structure"`

	// The availability options configured for the domain. Indicates whether Multi-AZ
	// is enabled for the domain.
	AvailabilityOptions *AvailabilityOptionsStatus `type:"structure"`
}

// String returns the string representation
func (s DescribeAvailabilityOptionsOutput) String() string {
	return awsutil.Prettify(s)
}