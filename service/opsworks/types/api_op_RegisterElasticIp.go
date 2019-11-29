// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RegisterElasticIpInput struct {
	_ struct{} `type:"structure"`

	// The Elastic IP address.
	//
	// ElasticIp is a required field
	ElasticIp *string `type:"string" required:"true"`

	// The stack ID.
	//
	// StackId is a required field
	StackId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RegisterElasticIpInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RegisterElasticIpInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RegisterElasticIpInput"}

	if s.ElasticIp == nil {
		invalidParams.Add(aws.NewErrParamRequired("ElasticIp"))
	}

	if s.StackId == nil {
		invalidParams.Add(aws.NewErrParamRequired("StackId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the response to a RegisterElasticIp request.
type RegisterElasticIpOutput struct {
	_ struct{} `type:"structure"`

	// The Elastic IP address.
	ElasticIp *string `type:"string"`
}

// String returns the string representation
func (s RegisterElasticIpOutput) String() string {
	return awsutil.Prettify(s)
}