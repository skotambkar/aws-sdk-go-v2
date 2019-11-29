// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetNetworkInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier of the network to get information about.
	//
	// NetworkId is a required field
	NetworkId *string `location:"uri" locationName:"networkId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetNetworkInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetNetworkInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetNetworkInput"}

	if s.NetworkId == nil {
		invalidParams.Add(aws.NewErrParamRequired("NetworkId"))
	}
	if s.NetworkId != nil && len(*s.NetworkId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NetworkId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetNetworkOutput struct {
	_ struct{} `type:"structure"`

	// An object containing network configuration parameters.
	Network *Network `type:"structure"`
}

// String returns the string representation
func (s GetNetworkOutput) String() string {
	return awsutil.Prettify(s)
}