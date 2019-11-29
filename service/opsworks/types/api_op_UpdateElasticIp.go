// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateElasticIpInput struct {
	_ struct{} `type:"structure"`

	// The IP address for which you want to update the name.
	//
	// ElasticIp is a required field
	ElasticIp *string `type:"string" required:"true"`

	// The new name.
	Name *string `type:"string"`
}

// String returns the string representation
func (s UpdateElasticIpInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateElasticIpInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateElasticIpInput"}

	if s.ElasticIp == nil {
		invalidParams.Add(aws.NewErrParamRequired("ElasticIp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateElasticIpOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateElasticIpOutput) String() string {
	return awsutil.Prettify(s)
}