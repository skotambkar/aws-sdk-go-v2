// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateEndpointInput struct {
	_ struct{} `type:"structure"`

	// The name of an endpoint configuration. For more information, see CreateEndpointConfig
	// (https://docs.aws.amazon.com/sagemaker/latest/dg/API_CreateEndpointConfig.html).
	//
	// EndpointConfigName is a required field
	EndpointConfigName *string `type:"string" required:"true"`

	// The name of the endpoint. The name must be unique within an AWS Region in
	// your AWS account.
	//
	// EndpointName is a required field
	EndpointName *string `type:"string" required:"true"`

	// An array of key-value pairs. For more information, see Using Cost Allocation
	// Tags (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html#allocation-what)in
	// the AWS Billing and Cost Management User Guide.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreateEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateEndpointInput"}

	if s.EndpointConfigName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointConfigName"))
	}

	if s.EndpointName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointName"))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateEndpointOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the endpoint.
	//
	// EndpointArn is a required field
	EndpointArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateEndpointOutput) String() string {
	return awsutil.Prettify(s)
}
