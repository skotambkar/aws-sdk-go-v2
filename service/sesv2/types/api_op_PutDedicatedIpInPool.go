// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to move a dedicated IP address to a dedicated IP pool.
type PutDedicatedIpInPoolInput struct {
	_ struct{} `type:"structure"`

	// The name of the IP pool that you want to add the dedicated IP address to.
	// You have to specify an IP pool that already exists.
	//
	// DestinationPoolName is a required field
	DestinationPoolName *string `type:"string" required:"true"`

	// The IP address that you want to move to the dedicated IP pool. The value
	// you specify has to be a dedicated IP address that's associated with your
	// AWS account.
	//
	// Ip is a required field
	Ip *string `location:"uri" locationName:"IP" type:"string" required:"true"`
}

// String returns the string representation
func (s PutDedicatedIpInPoolInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDedicatedIpInPoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDedicatedIpInPoolInput"}

	if s.DestinationPoolName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DestinationPoolName"))
	}

	if s.Ip == nil {
		invalidParams.Add(aws.NewErrParamRequired("Ip"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// An HTTP 200 response if the request succeeds, or an error message if the
// request fails.
type PutDedicatedIpInPoolOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDedicatedIpInPoolOutput) String() string {
	return awsutil.Prettify(s)
}
